package geo

import (
	"container/heap"
	"fmt"
	"log"
	"strconv"

	"github.com/neptunao/so-close/data"
)

func geoHeapTop(h *FixedSizeHeap, count int, center Coord) []RelativeCoord {
	res := make([]RelativeCoord, count)
	for i := 0; i < h.Limit; i++ {
		elem := heap.Pop(h).(Coord)
		res[i] = RelativeCoord{
			Coord:    elem,
			Center:   center,
			Distance: distance(center, elem),
		}
	}
	return res
}

// CalcTopPoints is a function to get TOP-(limit) nearest and furthest GeoPoints
// relative to center
func CalcTopPoints(center Coord, resultCount int, itr data.Iterator) (min []RelativeCoord, max []RelativeCoord, err error) {

	minHeap := MakeFixedSizeGeoDistMinHeap(MinPriorityQueue, resultCount, center)
	maxHeap := MakeFixedSizeGeoDistMinHeap(MaxPriorityQueue, resultCount, center)
	heap.Init(minHeap)
	heap.Init(maxHeap)
	i := 0
	for {
		data, ok := itr.Next()
		if !ok {
			break
		}
		record := data.([]string)
		lat, convErr := strconv.ParseFloat(record[1], 64)
		if convErr != nil {
			return nil, nil, convErr
		}
		lng, convErr := strconv.ParseFloat(record[2], 64)
		if convErr != nil {
			return nil, nil, convErr
		}
		coord := Coord{
			Name: record[0],
			Lat:  lat,
			Lon:  lng,
		}
		if !IsValidCoord(coord) {
			log.Printf("coordinate %d with value %s is invalid", i, coord)
			continue
		}
		heap.Push(minHeap, coord)
		heap.Push(maxHeap, coord)
		i++
	}
	if i < resultCount {
		return nil, nil, fmt.Errorf("wanted top %d but have only %d records",
			resultCount, i)
	}

	min = geoHeapTop(minHeap, resultCount, center)
	max = geoHeapTop(maxHeap, resultCount, center)
	return
}
