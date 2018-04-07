package main

// FixedSizeGeoDistMinHeap is a min binary heap implementation with distance
// from CenterCoord as priority. Size is fixed, new element always goes to the
// end and heapifies after insertion
type FixedSizeGeoDistMinHeap struct {
	data        []GeoCoord
	CenterCoord GeoCoord
	Limit       int
}

// MakeFixedSizeGeoDistMinHeap is a constructor for FixedSizeGeoDistMinHeap
func MakeFixedSizeGeoDistMinHeap(limit int, refCoord GeoCoord) *FixedSizeGeoDistMinHeap {
	return &FixedSizeGeoDistMinHeap{
		data:        make([]GeoCoord, 0, limit+1),
		CenterCoord: refCoord,
		Limit:       limit,
	}
}

func (h *FixedSizeGeoDistMinHeap) refDist(c GeoCoord) float64 {
	return distance(h.CenterCoord, c)
}

func (h *FixedSizeGeoDistMinHeap) Len() int {
	return len(h.data)
}

func (h *FixedSizeGeoDistMinHeap) Less(i, j int) bool {
	return h.refDist(h.data[i]) < h.refDist(h.data[j])
}

func (h *FixedSizeGeoDistMinHeap) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

// Push implementation of heap.Interface
func (h *FixedSizeGeoDistMinHeap) Push(x interface{}) {
	c := x.(GeoCoord)
	if len(h.data) <= h.Len() {
		h.data = append(h.data, c)
		return
	}
	h.data[h.Len()] = c
}

// Pop implementation of heap.Interface
func (h *FixedSizeGeoDistMinHeap) Pop() interface{} {
	n := len(h.data)
	x := h.data[n-1]
	h.data = h.data[0 : n-1]
	return x
}
