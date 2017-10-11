package main

type quickUnionFindRank struct {
	parent []int
	rank   []int
	count  int
}

func newQuickUnionFindRank(n int) *quickUnionFindRank {
	var quf = quickUnionFindRank{
		parent: make([]int, n),
		rank:   make([]int, n),
		count:  n,
	}
	for i := 0; i < n; i++ {
		quf.parent[i] = i
		quf.rank[i] = 1
	}

	return &quf
}

func (u *quickUnionFindRank) find(p int) int {
	if p < 0 || p >= u.count {
		panic("wrong index")
	}
	for p != u.parent[p] {
		p = u.parent[p]
	}
	return p
}

func (u *quickUnionFindRank) isConnected(p, q int) bool {
	return u.find(p) == u.find(q)
}

func (u *quickUnionFindRank) union(p, q int) {
	pID, qID := u.find(p), u.find(q)
	if pID == qID {
		return
	}

	if u.rank[pID] > u.rank[qID] {
		u.parent[qID] = pID
	} else if u.rank[qID] > u.rank[pID] {
		u.parent[pID] = qID
	} else {
		u.parent[qID] = pID
		u.rank[pID]++
	}
}
