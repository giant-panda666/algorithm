package main

type kruskalMST struct {
	g         graph
	mst       []*edge
	mstWeight Weight
}

func newKruskalMST(g graph) *kruskalMST {
	var k = kruskalMST{
		g:         g,
		mst:       make([]*edge, 0),
		mstWeight: IntWeight(0),
	}
	return &k
}

func (k *kruskalMST) kruskal() {
	var h = newMinHeap(k.g.E())
	for i := 0; i < k.g.V(); i++ {
		it := k.g.iter(i)
		for j := it.begin(); !it.end(); j = it.next() {
			if j.v() < j.w() {
				h.insert(j)

			}
		}
	}

	uf := newQuickUnionFindPathCompression(k.g.V())
	for !h.isEmpty() && len(k.mst) < k.g.V() {
		e := h.extractMin()
		if !uf.isConnected(e.v(), e.w()) {
			uf.union(e.v(), e.w())
			k.mst = append(k.mst, e)
		}
	}

	for i := 0; i < len(k.mst); i++ {
		k.mstWeight = k.mstWeight.Add(k.mst[i].wt())
	}
}

func (k *kruskalMST) result() Weight {
	return k.mstWeight
}
