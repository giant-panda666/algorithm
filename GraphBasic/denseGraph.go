package main

import "fmt"

type denseGraph struct {
	// n 节点个数，m 边条数
	n, m int
	// 是否是有向图
	directed bool
	// 使用邻接矩阵存储
	g [][]bool
}

func newDenseGraph(n int, directed bool) *denseGraph {
	var g = denseGraph{
		n:        n,
		directed: directed,
	}
	g.g = make([][]bool, n)
	for i := 0; i < n; i++ {
		g.g[i] = make([]bool, n)
	}
	return &g
}

func (g *denseGraph) V() int {
	return g.n
}

func (g *denseGraph) E() int {
	return g.m
}

func (g *denseGraph) addEdge(v, w int) {
	if v < 0 || v >= g.n || w < 0 || v >= g.n {
		panic("invalid edge")
	}
	if g.hasEdge(v, w) {
		return
	}

	g.g[v][w] = true
	if !g.directed {
		g.g[w][v] = true
	}
	g.m++
}

func (g *denseGraph) hasEdge(v, w int) bool {
	if v < 0 || v >= g.n || w < 0 || w >= g.n {
		panic("invalid edge")
	}
	return g.g[v][w]
}

func (g *denseGraph) show() {
	fmt.Println("+++++++++++denseGraph++++++++++")
	for i := 0; i < g.n; i++ {
		fmt.Println("node", i)
		iter := newDenseGraphIterator(g, i)
		for s := iter.begin(); !iter.end(); s = iter.next() {
			fmt.Print(s, " ")
		}
		fmt.Println()
	}
}

func (g *denseGraph) iter(v int) iterator {
	return newDenseGraphIterator(g, v)
}

func newDenseGraphIterator(dg *denseGraph, v int) *denseGraphIterator {
	var i = denseGraphIterator{
		dg: dg,
		v:  v,
	}
	return &i
}

type denseGraphIterator struct {
	dg    *denseGraph
	v     int
	index int
}

func (i *denseGraphIterator) begin() int {
	if i.dg.n == 0 {
		return -1
	}
	i.index = 0
	for ; i.index < i.dg.n; i.index++ {
		if i.dg.g[i.v][i.index] {
			return i.index
		}
	}
	return -1
}

func (i *denseGraphIterator) end() bool {
	return i.index >= i.dg.n
}

func (i *denseGraphIterator) next() int {
	for i.index++; i.index < i.dg.n; i.index++ {
		if i.dg.g[i.v][i.index] {
			return i.index
		}
	}
	return -1
}
