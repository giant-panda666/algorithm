package main

import "fmt"

type denseGraph struct {
	// n 节点个数，m 边条数
	n, m int
	// 是否是有向图
	directed bool
	// 使用邻接矩阵存储
	g [][]*edge
}

func newDenseGraph(n int, directed bool) *denseGraph {
	var g = denseGraph{
		n:        n,
		directed: directed,
	}
	g.g = make([][]*edge, n)
	for i := 0; i < n; i++ {
		g.g[i] = make([]*edge, n)
	}
	return &g
}

func (g *denseGraph) V() int {
	return g.n
}

func (g *denseGraph) E() int {
	return g.m
}

func (g *denseGraph) addEdge(v, w int, wt Weight) {
	if v < 0 || v >= g.n || w < 0 || v >= g.n {
		panic("invalid edge")
	}
	if g.hasEdge(v, w) {
		g.g[v][w] = nil
		if !g.directed {
			g.g[w][v] = nil
		}
		g.m--
	}

	g.g[v][w] = newEdge(v, w, wt)
	if !g.directed {
		g.g[w][v] = newEdge(w, v, wt)
	}
	g.m++
}

func (g *denseGraph) hasEdge(v, w int) bool {
	if v < 0 || v >= g.n || w < 0 || w >= g.n {
		panic("invalid edge")
	}
	return g.g[v][w] != nil
}

func (g *denseGraph) show() {
	fmt.Println("+++++++++++denseGraph++++++++++")
	for i := 0; i < g.n; i++ {
		for j := 0; j < g.n; j++ {
			if g.g[i][j] != nil {
				fmt.Printf("%v\t", g.g[i][j].wt())
			} else {
				fmt.Print("NULL\t")
			}
		}
		fmt.Println()
	}
}

func (g *denseGraph) iter(v int) iterator {
	return newDenseGraphIterator(g, v)
}

func newDenseGraphIterator(dg *denseGraph, v int) *denseGraphIterator {
	var i = denseGraphIterator{
		dg:    dg,
		v:     v,
		index: -1,
	}
	return &i
}

type denseGraphIterator struct {
	dg    *denseGraph
	v     int
	index int
}

func (i *denseGraphIterator) begin() *edge {
	if i.dg.n == 0 {
		return nil
	}
	i.index = -1
	return i.next()
}

func (i *denseGraphIterator) end() bool {
	return i.index >= i.dg.n
}

func (i *denseGraphIterator) next() *edge {
	for i.index++; i.index < i.dg.n; i.index++ {
		if i.dg.g[i.v][i.index] != nil {
			return i.dg.g[i.v][i.index]
		} else {
		}
	}
	return nil
}
