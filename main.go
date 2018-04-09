package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/neptunao/so-close/data"
	"github.com/neptunao/so-close/geo"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	centerLat = kingpin.Flag("center-lat", "Center geo point latitude").Default("51.925146").Float64()
	centerLon = kingpin.Flag("center-lon", "Center geo point longitude").Default("4.478617").Float64()
	filename  = kingpin.Arg("filename", "CSV file with geo data").Required().String()
)

func stringifyGeoCoordArray(coords []geo.RelativeCoord) []string {
	res := make([]string, len(coords))
	for i := 0; i < len(coords); i++ {
		res[i] = coords[i].String()
	}
	return res
}

// TODO read console arguments
func main() {
	kingpin.Parse()
	itr, err := data.ConnectCSVFile(*filename)
	if err != nil {
		log.Fatalf("error connecting to CSV file %s: %s", *filename, err)
	}
	defer itr.Close()
	const limit int = 5
	center := geo.Coord{
		Name: "HousingAnywhere Rotterdam office",
		Lat:  *centerLat,
		Lon:  *centerLon,
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
