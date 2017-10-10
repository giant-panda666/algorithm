package main

type quickUnionFind struct {
	parent []int
	count  int
}

func newQuickUnionFind(n int) *quickUnionFind {
	var quf = quickUnionFind{
		parent: make([]int, n),
		count:  n,
	}
	for i := 0; i < n; i++ {
		quf.parent[i] = i
	}
	return &quf
}

func (u *quickUnionFind) find(p int) int {
	if p < 0 || p >= u.count {
		panic("wrong index")
	}
	for p != u.parent[p] {
		p = u.parent[p]
	}
	return p
}

func (u *quickUnionFind) isConnected(p, q int) bool {
	return u.find(p) == u.find(q)
}

func (u *quickUnionFind) union(p, q int) {
	pID, qID := u.find(p), u.find(q)
	if pID == qID {
		return
	}

	u.parent[qID] = pID
}
