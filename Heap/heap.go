package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

type MaxHeap struct {
	item  []interface{}
	count int
}

func newMaxHeap(capacity int) *MaxHeap {
	var mh MaxHeap
	mh.item = make([]interface{}, 0, capacity+1)
	mh.count = 0
	return &mh
}

func (m *MaxHeap) size() int {
	return m.count
}

func (m *MaxHeap) isEmpty() bool {
	return m.count == 0
}
