package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/neptunao/so-close/data"
	"github.com/neptunao/so-close/geo"
)

func stringifyGeoCoordArray(coords []geo.Coord) []string {
	res := make([]string, len(coords))
	for i := 0; i < len(coords); i++ {
		res[i] = coords[i].String()
	}
	return res
}

// TODO read console arguments
func main() {
	const filename = "/home/neptunao/go/src/github.com/neptunao/so-close/geodata.csv"
	itr, err := data.ConnectCSVFile(filename)
	if err != nil {
		log.Fatalf("error connecting to CSV file %s: %s", filename, err)
	}
	defer itr.Close()
	const limit int = 5
	center := geo.Coord{
		Name: "HousingAnywhere Rotterdam office",
		Lat:  51.925146,
		Lon:  4.478617,
	}
	fmt.Printf("Calculating top %d nearest and furtherst points relative to %s\n",
		limit, center)
	min, max, err := geo.CalcTopPoints(center, limit, itr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Top %d nearest:\n", limit)
	fmt.Println(strings.Join(stringifyGeoCoordArray(min), "\n"))
	fmt.Printf("Top %d furtherst:\n", limit)
	fmt.Println(strings.Join(stringifyGeoCoordArray(max), "\n"))
}
