package main

import "fmt"

type bellmanFord struct {
	g               graph
	s               int
	distTo          []Weight
	from            []*edge
	hasNegtiveCycle bool
}

func newBellmanFord(g graph, s int) *bellmanFord {
	var b = bellmanFord{
		g:               g,
		s:               s,
		distTo:          make([]Weight, g.V()),
		from:            make([]*edge, g.V()),
		hasNegtiveCycle: false,
	}
	b.distTo[s] = IntWeight(0)
	return &b
}

func (b *bellmanFord) BellmanFord() {
	for pass := 0; pass < b.g.V(); pass++ {
		for i := 0; i < b.g.V(); i++ {
			it := b.g.iter(i)
			for j := it.begin(); !it.end(); j = it.next() {
				if b.distTo[j.w()] == nil || b.distTo[i].Add(j.wt()).Less(b.distTo[j.w()]) {
					b.distTo[j.w()] = b.distTo[i].Add(j.wt())
					b.from[j.w()] = j
				}
			}
		}
	}

	b.detectNegtiveCycle()
}

func (b *bellmanFord) detectNegtiveCycle() {
	for i := 0; i < b.g.V(); i++ {
		it := b.g.iter(i)
		for j := it.begin(); !it.end(); j = it.next() {
			if b.distTo[j.w()] == nil || b.distTo[i].Add(j.wt()).Less(b.distTo[j.w()]) {
				b.hasNegtiveCycle = true
			}
		}
	}
}

func (b *bellmanFord) shortestPath(w int) []*edge {
	if w < 0 || w >= b.g.V() {
		return nil
	}

	var ret []*edge
	e := b.from[w]
	ret = append(ret, e)
	for e.v() != b.s {
		e = b.from[e.v()]
		ret = append(ret, e)
		ret = append(ret[len(ret)-1:], ret[:len(ret)-1]...)
	}

	return ret
}

func (b *bellmanFord) showPath(w int) {
	if w < 0 || w >= b.g.V() {
		return
	}
	path := b.shortestPath(w)
	for i, v := range path {
		fmt.Print(v.v(), "->")
		if i == len(path)-1 {
			fmt.Println(v.w())
		}
	}
}
