package main

type graph interface {
	V() int
	E() int
	addEdge(v, w int)
	hasEdge(v, w int) bool
	show()
	iter(v int) iterator
}

type iterator interface {
	begin() int
	end() bool
	next() int
}
