package main

type quickUnionFindSize struct {
	parent []int
	sz     []int
	count  int
}

func newQuickUnionFindSize(n int) *quickUnionFindSize {
	var quf = quickUnionFindSize{
		parent: make([]int, n),
		sz:     make([]int, n),
		count:  n,
	}
	for i := 0; i < n; i++ {
		quf.parent[i] = i
		quf.sz[i] = 1
	}
	return &quf
}

func (u *quickUnionFindSize) find(p int) int {
	if p < 0 || p >= u.count {
		panic("wrong index")
	}
	for p != u.parent[p] {
		p = u.parent[p]
	}
	return p
}

func (u *quickUnionFindSize) isConnected(p, q int) bool {
	return u.find(p) == u.find(q)
}

func (u *quickUnionFindSize) union(p, q int) {
	pID, qID := u.find(p), u.find(q)
	if pID == qID {
		return
	}

	if u.sz[pID] > u.sz[qID] {
		u.parent[pID] = qID
		u.sz[pID] += u.sz[qID]
	} else {
		u.parent[qID] = pID
		u.sz[qID] += u.sz[pID]
	}
}
