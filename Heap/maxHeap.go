package main

type MaxHeap struct {
	item     []int
	count    int
	capacity int
}

func newMaxHeap(capacity int) *MaxHeap {
	var mh MaxHeap
	mh.item = make([]int, 1, capacity+1)
	mh.count = 0
	mh.capacity = capacity
	return &mh
}

func (m *MaxHeap) size() int {
	return m.count
}

func (m *MaxHeap) isEmpty() bool {
	return m.count == 0
}

func (m *MaxHeap) shiftUp(k int) {
	for k > 1 && m.item[k] > m.item[k/2] {
		m.item[k], m.item[k/2] = m.item[k/2], m.item[k]
		k /= 2
	}
}

func (m *MaxHeap) shiftDown(k int) {
	for 2*k <= m.count {
		j := 2 * k
		if j+1 < m.count && m.item[j+1] > m.item[j] {
			j++
		}
		if m.item[j] < m.item[k] {
			return
		}
		m.item[j], m.item[k] = m.item[k], m.item[j]
		k = j
	}
}

func (m *MaxHeap) Insert(data int) {
	if m.count+1 > m.capacity {
		panic("full data")
	}

	m.item = append(m.item, data)
	m.shiftUp(m.count + 1)
	m.count++
}

func (m *MaxHeap) ExtractMax() int {
	if m.count <= 0 {
		panic("no data")
	}

	ret := m.item[1]
	m.item[1] = m.item[m.count]
	m.count--
	m.shiftDown(1)

	return ret
}

func (m *MaxHeap) GetMax() int {
	if m.count <= 0 {
		panic("no data")
	}

	return m.item[1]
}
