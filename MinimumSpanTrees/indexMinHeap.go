package main

type indexMinHeap struct {
	// 数据
	data []*edge
	// 索引堆，可以通过indexes[i]得到数据中的索引
	indexes []int
	// 数据到堆的反向索引，可以通过reverse[i]得到数据data[i]在堆中的位置
	reverse  []int
	count    int
	capacity int
}

func newIndexMinHeap(capacity int) *indexMinHeap {
	var m = indexMinHeap{
		data:     make([]*edge, 1+capacity),
		indexes:  make([]int, 1+capacity),
		reverse:  make([]int, 1+capacity),
		count:    0,
		capacity: capacity,
	}

	return &m
}

func (h *indexMinHeap) shiftUP(k int) {
	for k > 1 && h.data[h.indexes[k]].less(h.data[h.indexes[k/2]]) {
		h.indexes[k], h.indexes[k/2] = h.indexes[k/2], h.indexes[k]
		h.reverse[h.indexes[k]], h.reverse[h.indexes[k/2]] = k, k/2
		k /= 2
	}
}

func (h *indexMinHeap) shiftDOWN(k int) {
	for 2*k <= h.count {
		j := 2 * k
		if j+1 <= h.count && h.data[h.indexes[j+1]].less(h.data[h.indexes[j]]) {
			j += 1
		}
		if h.data[h.indexes[k]].less(h.data[h.indexes[j]]) {
			break
		}
		h.indexes[j], h.indexes[k] = h.indexes[k], h.indexes[j]
		h.reverse[h.indexes[j]], h.reverse[h.indexes[k]] = j, k
		k = j
	}
}

func (h *indexMinHeap) extractMin() *edge {
	if h.count <= 0 {
		return nil
	}
	ret := h.data[h.indexes[1]]
	h.indexes[1], h.indexes[h.count] = h.indexes[h.count], h.indexes[1]
	h.reverse[h.indexes[1]], h.reverse[h.indexes[h.count]] = 1, 0
	h.count--
	h.shiftDOWN(1)
	return ret
}

func (h *indexMinHeap) extractMinIndex() int {
	if h.count <= 0 {
		return -1
	}
	index := h.indexes[1] - 1
	h.indexes[1], h.indexes[h.count] = h.indexes[h.count], h.indexes[1]
	h.reverse[h.indexes[1]], h.reverse[h.indexes[h.count]] = 1, 0
	h.count--
	h.shiftDOWN(1)
	return index
}

func (h *indexMinHeap) contain(i int) bool {
	if i < 0 || i > h.capacity {
		return false
	}
	return h.reverse[i+1] != 0
}

func (h *indexMinHeap) getEdge(i int) *edge {
	if !h.contain(i) {
		return nil
	}
	return h.data[i+1]
}

func (h *indexMinHeap) change(i int, e *edge) {
	if !h.contain(i) {
		return
	}
	i++
	h.data[i] = e
	// 找到indexes[j] = i, j表示data[i]在堆中的位置
	// 之后进行shiftUP，shiftDOWN操作
	//	for j := 1; j <= h.count; j++ {
	//		if h.indexes[j] == i {
	//			h.shiftUP(j)
	//			h.shiftDOWN(j)
	//			return
	//		}
	//	}
	j := h.reverse[i]
	h.shiftUP(j)
	h.shiftDOWN(j)
}

func (h *indexMinHeap) insert(i int, data *edge) {
	if h.count+1 > h.capacity {
		return
	}
	if i < 0 || i+1 > h.capacity {
		return
	}
	i++
	h.data[i] = data
	h.indexes[h.count+1] = i
	h.reverse[i] = h.count + 1
	h.shiftUP(h.count + 1)
	h.count++
}

func (h *indexMinHeap) isEmpty() bool {
	return h.count == 0
}

func (h *indexMinHeap) size() int {
	return h.count
}

func (h *indexMinHeap) getMin() *edge {
	if h.count <= 0 {
		return nil
	}
	return h.data[h.indexes[1]]
}
