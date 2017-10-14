package main

type quickUnionFindPathCompression struct {
	parent []int
	rank   []int
	count  int
}

func newQuickUnionFindPathCompression(n int) *quickUnionFindPathCompression {
	var quf = quickUnionFindPathCompression{
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

func (u *quickUnionFindPathCompression) find(p int) int {
	if p < 0 || p >= u.count {
		panic("wrong index")
	}
	//	for p != u.parent[p] {
	//		// 路径压缩，让p指向父亲节点的父亲
	//		u.parent[p] = u.parent[u.parent[p]]
	//		p = u.parent[p]
	//	}
	//  return p

	if p != u.parent[p] {
		// 更彻底的压缩
		u.parent[p] = u.find(u.parent[p])
	}
	return u.parent[p]
}

func (u *quickUnionFindPathCompression) isConnected(p, q int) bool {
	return u.find(p) == u.find(q)
}

func (u *quickUnionFindPathCompression) union(p, q int) {
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
