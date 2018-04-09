package geo

// PriorityQueueMode is a type to specify PriorityQueue order: minimum or maximum
type PriorityQueueMode int

const (
	// MinPriorityQueue maps to Minimal Priority Queue (pops element with minimal priority first)
	MinPriorityQueue PriorityQueueMode = iota
	// MaxPriorityQueue maps to Maximum Priority Queue (pops element with maximal priority first)
	MaxPriorityQueue
)

// FixedSizeHeap is a min binary heap implementation with distance
// from CenterCoord as priority. Size is fixed, new element always goes to the
// end and heapifies after insertion
type FixedSizeHeap struct {
	data        []Coord
	CenterCoord Coord
	Limit       int
	mode        PriorityQueueMode
}

// MakeFixedSizeGeoDistMinHeap is a constructor for FixedSizeHeap
func MakeFixedSizeGeoDistMinHeap(mode PriorityQueueMode, limit int,
	refCoord Coord) *FixedSizeHeap {

	return &FixedSizeHeap{
		data:        make([]Coord, 0, limit+1),
		CenterCoord: refCoord,
		Limit:       limit,
		mode:        mode,
	}
}

func (h *FixedSizeHeap) refDist(c Coord) float64 {
	return distance(h.CenterCoord, c)
}

func (h *FixedSizeHeap) Len() int {
	return len(h.data)
}

func (h *FixedSizeHeap) Less(i, j int) bool {
	switch h.mode {
	case MinPriorityQueue:
		return h.refDist(h.data[i]) < h.refDist(h.data[j])
	case MaxPriorityQueue:
		return h.refDist(h.data[i]) > h.refDist(h.data[j])
	default:
		return h.refDist(h.data[i]) < h.refDist(h.data[j])
	}
}

func (h *FixedSizeHeap) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

// Push implementation of heap.Interface
func (h *FixedSizeHeap) Push(x interface{}) {
	c := x.(Coord)
	if len(h.data) <= h.Len() {
		h.data = append(h.data, c)
		return
	}
	h.data[h.Len()] = c
}

// Pop implementation of heap.Interface
func (h *FixedSizeHeap) Pop() interface{} {
	n := len(h.data)
	x := h.data[n-1]
	h.data = h.data[0 : n-1]
	return x
}
