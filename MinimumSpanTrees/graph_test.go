package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestGraph(t *testing.T) {
	N, M := 6, 10
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
		wt := rand.Int() % 100
		g.addEdge(a, b, IntWeight(wt))
	}
	g.show()
}
