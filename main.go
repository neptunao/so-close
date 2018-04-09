package main

import (
	"container/heap"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/neptunao/so-close/data"
)

func geoHeapTop(h *FixedSizeGeoDistMinHeap, count int) []GeoCoord {
	res := make([]GeoCoord, count)
	for i := 0; i < h.Limit; i++ {
		elem := heap.Pop(h).(GeoCoord)
		res[i] = elem
	}
	return res
}

func stringifyGeoCoordArray(coords []GeoCoord) []string {
	res := make([]string, len(coords))
	for i := 0; i < len(coords); i++ {
		res[i] = coords[i].String()
	}
	return res
}

// CalcTopPoints is a function to get TOP-(limit) nearest and furtherst GeoPoints
// relative to center
func CalcTopPoints(center GeoCoord, resultCount int, itr data.Iterator) (min []GeoCoord, max []GeoCoord, err error) {
	minHeap := MakeFixedSizeGeoDistMinHeap(MinPriorityQueue, resultCount, center)
	maxHeap := MakeFixedSizeGeoDistMinHeap(MaxPriorityQueue, resultCount, center)
	heap.Init(minHeap)
	heap.Init(maxHeap)

	itr.Next() // Skip header with field names
	for {
		data, ok := itr.Next()
		if !ok {
			break
		}
		record := data.([]string)
		latStr, convErr := strconv.ParseFloat(record[1], 64)
		if convErr != nil {
			return nil, nil, convErr
		}
		lngStr, convErr := strconv.ParseFloat(record[2], 64)
		if convErr != nil {
			return nil, nil, convErr
		}
		coord := GeoCoord{
			Name: record[0],
			Lat:  latStr,
			Lon:  lngStr,
		}
		heap.Push(minHeap, coord)
		heap.Push(maxHeap, coord)
	}

	min = geoHeapTop(minHeap, resultCount)
	max = geoHeapTop(maxHeap, resultCount)
	return
}

func main() {
	const filename = "/home/neptunao/go/src/github.com/neptunao/so-close/geodata.csv"
	itr, err := data.ConnectCSVFile(filename)
	if err != nil {
		log.Fatalf("error connecting to CSV file %s: %s", filename, err)
	}
	defer itr.Close()
	const limit int = 5
	center := GeoCoord{
		Name: "HousingAnywhere Rotterdam office",
		Lat:  51.925146,
		Lon:  4.478617,
	}
	fmt.Printf("Calculating top %d nearest and furtherst points relative to %s\n",
		limit, center)
	min, max, err := CalcTopPoints(center, limit, itr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Top %d nearest:\n", limit)
	fmt.Println(strings.Join(stringifyGeoCoordArray(min), "\n"))
	fmt.Printf("Top %d furtherst:\n", limit)
	fmt.Println(strings.Join(stringifyGeoCoordArray(max), "\n"))
}
