package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestGraph(t *testing.T) {
	N, M := 3000, 10000
	dg := newDenseGraph(N, false)
	sg := newSparseGraph(N, false)
	fmt.Println("denseGraph")
	showGraph(dg, N, M)
	fmt.Println("denseGraph Prim")
	showPrim(dg)
	fmt.Println("=============================================================")
	fmt.Println("sparseGraph")
	showGraph(sg, N, M)
	fmt.Println("sparseGraph Prim")
	showPrim(sg)
}

func showGraph(g graph, N, M int) {
	rand.Seed(time.Now().Unix())
	// prim算法要求图必须是连通的，因此这里生成的图需要是连通的
	for i := 1; i < N; i++ {
		wt := rand.Int() % 100
		g.addEdge(0, i, IntWeight(wt))
	}
	for i := 0; i < M-N+1; i++ {
		a := rand.Int() % N
		b := rand.Int() % N
		wt := rand.Int() % 100
		if a == b {
			continue
		}
		g.addEdge(a, b, IntWeight(wt))
	}
	//	g.addEdge(0, 1, IntWeight(30))
	//	g.addEdge(0, 2, IntWeight(77))
	//	g.addEdge(0, 3, IntWeight(23))
	//	g.addEdge(0, 4, IntWeight(11))
	//	g.addEdge(1, 3, IntWeight(68))
	//	g.addEdge(2, 4, IntWeight(15))
	//	g.addEdge(3, 4, IntWeight(97))
	//	g.show()
}

func showPrim(g graph) {
	lazyPrimMST := newLazyPrimMST(g)
	start := time.Now()
	lazyPrimMST.Prim()
	fmt.Println("lazyPrimMST using", time.Since(start))
	//	edges1 := lazyPrimMST.mstEdges()
	//	for _, v := range edges1 {
	//		fmt.Println(v.v(), "to", v.w())
	//	}
	fmt.Println("lazyprim mstWeight", lazyPrimMST.result())

	primMST := newPrimMST(g)
	start = time.Now()
	primMST.Prim()
	fmt.Println("primMST using", time.Since(start))
	//	edges2 := primMST.mstEdges()
	//	for _, v := range edges2 {
	//		fmt.Println(v.v(), "to", v.w())
	//	}
	fmt.Println("prim mstWeight", primMST.result())
}
