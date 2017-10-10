package main

type unionFind struct {
	id    []int
	count int
}

func newUnionFind(n int) *unionFind {
	var uf = unionFind{
		id:    make([]int, n),
		count: n,
	}
	for i := 0; i < n; i++ {
		uf.id[i] = i
	}
	return &uf
}

func (u *unionFind) find(p int) int {
	if p < 0 || p >= u.count {
		panic("wrong index")
	}
	return u.id[p]
}

func (u *unionFind) isConnected(p, q int) bool {
	return u.find(p) == u.find(q)
}

func (u *unionFind) union(p, q int) {
	pID, qID := u.find(p), u.find(q)
	if pID == qID {
		return
	}

	for i := 0; i < u.count; i++ {
		if u.find(i) == pID {
			u.id[i] = qID
		}
	}
}
