package main

import (
	"container/heap"
	"encoding/csv"
	"io"
	"os"
	"strconv"
)

func CalcTopDists(center GeoCoord, resultCount int) (min []GeoCoord, max []GeoCoord, err error) {
	//TODO make an abstraction
	f, err := os.Open("/home/neptunao/Downloads/geoData.csv")
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()
	r := csv.NewReader(f)
	minHeap := MakeFixedSizeGeoDistMinHeap(resultCount, center)
	heap.Init(minHeap)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, nil, err
		}

		//TODO check errors
		latStr, _ := strconv.ParseFloat(record[1], 64)
		lngStr, _ := strconv.ParseFloat(record[2], 64)
		heap.Push(minHeap, GeoCoord{
			Name: record[0],
			Lat:  latStr,
			Lon:  lngStr,
		})
	}
	min = make([]GeoCoord, resultCount)
	max = make([]GeoCoord, resultCount)
	for i := 0; i < minHeap.Limit; i++ {
		elem := heap.Pop(minHeap).(GeoCoord)
		// fmt.Printf("%d: %v distance=%f\n", i, elem, distance(center, elem))
		min[i] = elem
	}
	return min, max, nil
}

func main() {
	CalcTopDists(GeoCoord{Lat: 51.925146, Lon: 4.478617}, 5)
}
