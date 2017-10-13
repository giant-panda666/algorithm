package main

type minHeap struct {
	data     []*edge
	count    int
	capacity int
}

func newMinHeap(capacity int) *minHeap {
	var m = minHeap{
		data:     make([]*edge, 1+capacity),
		count:    0,
		capacity: capacity,
	}

	return &m
}

func (h *minHeap) shiftUP(k int) {
	for k > 1 && h.data[k].less(h.data[k/2]) {
		h.data[k], h.data[k/2] = h.data[k/2], h.data[k]
		k /= 2
	}
}

func (h *minHeap) shiftDOWN(k int) {
	for 2*k <= h.count {
		j := 2 * k
		if j+1 <= h.count && h.data[j+1].less(h.data[j]) {
			j += 1
		}
		if h.data[k].less(h.data[j]) {
			break
		}
		h.data[j], h.data[k] = h.data[k], h.data[j]
		k = j
	}
}

func (h *minHeap) extractMin() *edge {
	if h.count <= 0 {
		return nil
	}
	ret := h.data[1]
	h.data[1], h.data[h.count] = h.data[h.count], h.data[1]
	h.count--
	h.shiftDOWN(1)
	return ret
}

func (h *minHeap) insert(data *edge) {
	if h.count+1 > h.capacity {
		return
	}
	h.data[h.count+1] = data
	h.shiftUP(h.count + 1)
	h.count++
}

func (h *minHeap) isEmpty() bool {
	return h.count == 0
}

func (h *minHeap) size() int {
	return h.count
}

func (h *minHeap) getMin() *edge {
	if h.count <= 0 {
		return nil
	}
	return h.data[1]
}
