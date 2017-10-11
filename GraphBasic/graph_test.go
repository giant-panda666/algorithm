package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestGraph(t *testing.T) {
	N, M := 20, 30
	dg := newDenseGraph(N, false)
	sg := newSparseGraph(N, false)
	fmt.Println("denseGraph")
	showGraph(dg, N, M)
	fmt.Println("sparseGraph")
	showGraph(sg, N, M)
}

func showGraph(g graph, N, M int) {
	rand.Seed(time.Now().Unix())
	for i := 0; i < M; i++ {
		a := rand.Int() % N
		b := rand.Int() % N
		g.addEdge(a, b)
	}
	g.show()

	c1 := newComponent(g)
	c1.DFS()
	p1 := newPath(g, 0)
	p1.showPath(rand.Int() % N)
	p11 := newShortestPath(g, 0)
	p11.showPath(rand.Int() % N)

	c2 := newComponent(g)
	c2.BFS()
	p2 := newPath(g, 0)
	p2.showPath(rand.Int() % N)
	p22 := newShortestPath(g, 0)
	p22.showPath(rand.Int() % N)
}
