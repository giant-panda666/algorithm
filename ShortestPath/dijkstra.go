package main

import (
	"fmt"
)

type dijkstra struct {
	g      graph
	s      int
	distTo []Weight
	marked []bool
	from   []*edge
}

func newDijkstra(g graph, s int) *dijkstra {
	var d = dijkstra{
		g:      g,
		s:      s,
		distTo: make([]Weight, g.V()),
		marked: make([]bool, g.V()),
		from:   make([]*edge, g.V()),
	}
	return &d
}

func (d *dijkstra) Dijkstra() {
	ipq := newIndexMinHeap(d.g.V())
	d.distTo[d.s] = IntWeight(0)
	ipq.insert(d.s, d.distTo[d.s])

	// 松弛操作。每次从最小堆中取出节点v到最短路径点集合中，然后查看与该点相连的点，
	// 如果当前源点s到该点的最短距离大于通过v到达该点的距离，则源点s到该点的最短距离
	// 更新（插入）到最小堆中。
	for !ipq.isEmpty() {
		v := ipq.extractMinIndex()
		d.marked[v] = true

		i := d.g.iter(v)
		for j := i.begin(); !i.end(); j = i.next() {
			w := j.other(v)
			if !d.marked[w] {
				if d.distTo[w] == nil || d.distTo[v].Add(j.wt()).Less(d.distTo[w]) {
					d.distTo[w] = d.distTo[v].Add(j.wt())
					d.from[w] = j
					if ipq.contain(w) {
						ipq.change(w, d.distTo[w])
					} else {
						ipq.insert(w, d.distTo[w])
					}
				}
			}
		}
	}
}

func (d *dijkstra) shortestPathTo(w int) Weight {
	if w < 0 || w >= d.g.V() {
		return nil
	}
	return d.distTo[w]
}

func (d *dijkstra) hasPathTo(w int) bool {
	if w < 0 || w >= d.g.V() {
		return false
	}
	return d.marked[w]
}

func (d *dijkstra) shortestPath(w int) []*edge {
	if w < 0 || w >= d.g.V() {
		return nil
	}

	var ret []*edge
	e := d.from[w]
	ret = append(ret, e)
	for e.v() != d.s {
		e = d.from[e.v()]
		ret = append(ret, e)
		ret = append(ret[len(ret)-1:], ret[:len(ret)-1]...)
	}

	return ret
}

func (d *dijkstra) showPath(w int) {
	if w < 0 || w >= d.g.V() {
		return
	}
	path := d.shortestPath(w)
	for i, v := range path {
		fmt.Print(v.v(), "->")
		if i == len(path)-1 {
			fmt.Println(v.w())
		}
	}
}
