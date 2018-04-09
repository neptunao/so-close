package geo

import (
	"container/heap"
	"strconv"

	"github.com/neptunao/so-close/data"
)

func geoHeapTop(h *FixedSizeHeap, count int) []Coord {
	res := make([]Coord, count)
	for i := 0; i < h.Limit; i++ {
		elem := heap.Pop(h).(Coord)
		res[i] = elem
	}
	return res
}

// CalcTopPoints is a function to get TOP-(limit) nearest and furtherst GeoPoints
// relative to center
func CalcTopPoints(center Coord, resultCount int, itr data.Iterator) (min []Coord, max []Coord, err error) {
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
		coord := Coord{
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
