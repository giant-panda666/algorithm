package main

import "fmt"

type sparseGraph struct {
	n, m     int
	directed bool
	g        [][]*edge
}

func newSparseGraph(n int, directed bool) *sparseGraph {
	var g = sparseGraph{
		n:        n,
		directed: directed,
		g:        make([][]*edge, n),
	}
	return &g
}

func (g *sparseGraph) V() int {
	return g.n
}

func (g *sparseGraph) E() int {
	return g.m
}

func (g *sparseGraph) addEdge(v, w int, wt Weight) {
	if v < 0 || v >= g.n || w < 0 || w >= g.n {
		panic("invalid edge")
	}
	if g.hasEdge(v, w) {
		return
	}

	g.g[v] = append(g.g[v], newEdge(v, w, wt))
	if !g.directed {
		g.g[w] = append(g.g[w], newEdge(w, v, wt))
	}
	g.m++
}

func (g *sparseGraph) hasEdge(v, w int) bool {
	if v < 0 || v >= g.n || w < 0 || w >= g.n {
		panic("invalid edge")
	}
	for i := 0; i < len(g.g[v]); i++ {
		if g.g[v][i].other(v) == w {
			return true
		}
	}
	return false
}

func (g *sparseGraph) show() {
	fmt.Println("+++++++++++sparseGraph++++++++++")
	for i := 0; i < g.n; i++ {
		fmt.Println("node", i)
		for j := 0; j < len(g.g[i]); j++ {
			if j != 0 && j%8 == 0 {
				fmt.Println()
			}
			fmt.Printf("(to %d, wt:%d)\t", g.g[i][j].w(), g.g[i][j].wt())
		}
		fmt.Println()
	}
}

func (g *sparseGraph) iter(v int) iterator {
	return newSparseGraphIterator(g, v)
}

func newSparseGraphIterator(sg *sparseGraph, v int) *sparseGraphIterator {
	var i = sparseGraphIterator{
		sg: sg,
		v:  v,
	}
	return &i
}

type sparseGraphIterator struct {
	sg    *sparseGraph
	v     int
	index int
}

func (i *sparseGraphIterator) begin() *edge {
	if i.sg.n == 0 {
		return nil
	}
	i.index = 0
	if len(i.sg.g[i.v]) != 0 {
		return i.sg.g[i.v][0]
	}
	return nil
}

func (i *sparseGraphIterator) end() bool {
	return i.index >= len(i.sg.g[i.v])
}

func (i *sparseGraphIterator) next() *edge {
	i.index++
	if len(i.sg.g[i.v]) > i.index {
		return i.sg.g[i.v][i.index]
	}
	return nil
}
