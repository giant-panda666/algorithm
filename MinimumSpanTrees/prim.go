package main

type primMST struct {
	g graph
	// 横切边集合，每次从横切边中选择最短的边，因此使用最小堆保存
	ipq *indexMinHeap
	// 区分已选择和未选择的节点
	marked []bool
	// 最小生成树
	mst []*edge
	// 总权重
	mstWeight Weight
}

func newPrimMST(g graph) *primMST {
	var t = primMST{
		g:         g,
		ipq:       newIndexMinHeap(g.V()),
		marked:    make([]bool, g.V()),
		mst:       make([]*edge, 0),
		mstWeight: IntWeight(0),
	}

	return &t
}

func (t *primMST) Prim() {
	// 从0节点开始生成
	t.visit(0)
	for !t.ipq.isEmpty() {
		e := t.ipq.extractMin()
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

func (t *primMST) visit(v int) {
	t.marked[v] = true

	i := t.g.iter(v)
	for j := i.begin(); !i.end(); j = i.next() {
		if j != nil && !t.marked[j.other(v)] {
			tmp := t.ipq.getEdge(j.other(v))
			if tmp == nil {
				t.ipq.insert(j.other(v), j)
			} else {
				if j.less(tmp) {
					t.ipq.change(j.other(v), j)
				}
			}
		}
	}
}

func (t *primMST) result() Weight {
	return t.mstWeight
}

func (t *primMST) mstEdges() []*edge {
	return t.mst
}
