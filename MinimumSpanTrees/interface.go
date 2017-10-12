package main

type graph interface {
	V() int
	E() int
	addEdge(v, w int, weight Weight)
	hasEdge(v, w int) bool
	show()
	iter(v int) iterator
}

type iterator interface {
	begin() *edge
	end() bool
	next() *edge
}
