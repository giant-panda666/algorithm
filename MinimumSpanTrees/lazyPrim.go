package main

type lazyPrimMST struct {
	g graph
	// 横切边集合，每次从横切边中选择最短的边，因此使用最小堆保存
	pq *minHeap
	// 区分已选择和未选择的节点
	marked []bool
	// 最小生成树
	mst []*edge
	// 总权重
	mstWeight Weight
}

func newLazyPrimMST(g graph) *lazyPrimMST {
	var t = lazyPrimMST{
		g:         g,
		pq:        newMinHeap(g.E()),
		marked:    make([]bool, g.V()),
		mst:       make([]*edge, 0),
		mstWeight: IntWeight(0),
	}

	return &t
}

func (t *lazyPrimMST) Prim() {
	// 从0节点开始生成
	t.visit(0)
	for !t.pq.isEmpty() {
		e := t.pq.extractMin()
		if t.marked[e.v()] == t.marked[e.w()] {
			continue
		}

		t.mst = append(t.mst, e)
		if !t.marked[e.v()] {
			t.visit(e.v())
		} else {
			t.visit(e.w())
		}
	}

	for i := 0; i < len(t.mst); i++ {
		t.mstWeight = t.mstWeight.Add(t.mst[i].wt())
	}
}

func (t *lazyPrimMST) visit(v int) {
	t.marked[v] = true

	i := t.g.iter(v)
	for j := i.begin(); !i.end(); j = i.next() {
		if j != nil && !t.marked[j.other(v)] {
			t.pq.insert(j)
		}
	}
}

func (t *lazyPrimMST) result() Weight {
	return t.mstWeight
}

func (t *lazyPrimMST) mstEdges() []*edge {
	return t.mst
}
