package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestGraph(t *testing.T) {
	N, M := 10, 20
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
		g.addEdge(a, b, IntWeight(wt))
	}
	g.show()
}

func showPrim(g graph) {
	primMST := newLazyPrimMST(g)
	primMST.Prim()
	edges := primMST.mstEdges()
	for _, v := range edges {
		fmt.Println(v.v(), "to", v.w())
	}
}
