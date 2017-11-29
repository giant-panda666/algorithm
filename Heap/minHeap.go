package main

import "fmt"

func main() {
	var mh = newMinHeap(10)
	for i := 0; i < 10; i++ {
		mh.Insert(10 - i)
	}
	for !mh.isEmpty() {
		fmt.Print(mh.ExtractMin(), " ")
	}
	fmt.Println()
}

type MinHeap struct {
	item     []int
	count    int
	capacity int
}

func newMinHeap(capacity int) *MinHeap {
	var mh MinHeap
	mh.item = make([]int, capacity+1)
	mh.count = 0
	mh.capacity = capacity
	return &mh
}

func (m *MinHeap) size() int {
	return m.count
}

func (m *MinHeap) isEmpty() bool {
	return m.count == 0
}

func (m *MinHeap) shiftUp(i int) {
	for i/2 >= 1 && m.item[i] < m.item[i/2] {
		m.item[i], m.item[i/2] = m.item[i/2], m.item[i]
		i = i / 2
	}
}

func (m *MinHeap) Insert(data int) {
	if m.count == m.capacity {
		panic("no space")
	}
	m.count++
	m.item[m.count] = data
	m.shiftUp(m.count)
}

func (m *MinHeap) shiftDown(i int) {
	for 2*i <= m.count {
		min := 2 * i
		if 2*i+1 <= m.count && m.item[2*i+1] < m.item[min] {
			min = 2*i + 1
		}
		if m.item[i] < m.item[min] {
			return
		}

		m.item[i], m.item[min] = m.item[min], m.item[i]
		i = min
	}
}

func (m *MinHeap) ExtractMin() int {
	if m.count <= 0 {
		panic("no data")
	}

	ret := m.item[1]
	m.item[1] = m.item[m.count]
	m.count--
	m.shiftDown(1)
	return ret
}

func (m *MinHeap) GetMin() int {
	if m.count < 0 {
		panic("no data")
	}

	return m.item[1]
}
