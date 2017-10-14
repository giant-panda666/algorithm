package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestGraph(t *testing.T) {
	N, M := 5, 10
	dg := newDenseGraph(N, true)
	sg := newSparseGraph(N, true)
	fmt.Println("denseGraph")
	showGraph(dg, N, M)
	showDijkstra(dg)
	showBellmanFord(dg)
	fmt.Println("=============================================================")
	fmt.Println("sparseGraph")
	showGraph(sg, N, M)
	showDijkstra(sg)
	showBellmanFord(sg)
}

func showGraph(g graph, N, M int) {
	rand.Seed(time.Now().Unix())
	//	for i := 0; i < M-N+1; i++ {
	//		a := rand.Int() % N
	//		b := rand.Int() % N
	//		wt := rand.Int() % 100
	//		if a == b {
	//			i--
	//			continue
	//		}
	//		g.addEdge(a, b, IntWeight(wt))
	//	}
	g.addEdge(0, 1, IntWeight(30))
	g.addEdge(0, 2, IntWeight(77))
	g.addEdge(0, 3, IntWeight(23))
	g.addEdge(0, 4, IntWeight(100))
	g.addEdge(1, 3, IntWeight(68))
	g.addEdge(2, 4, IntWeight(15))
	g.addEdge(3, 4, IntWeight(97))
	g.show()
}

func showDijkstra(g graph) {
	fmt.Println("dijkstra")
	d := newDijkstra(g, 0)
	d.Dijkstra()
	for i := 1; i < g.V(); i++ {
		fmt.Println("show path to", i)
		d.showPath(i)
	}
}

func showBellmanFord(g graph) {
	fmt.Println("bellmanFord")
	b := newBellmanFord(g, 0)
	b.BellmanFord()
	for i := 1; i < g.V(); i++ {
		fmt.Println("show path to", i)
		b.showPath(i)
	}
}
