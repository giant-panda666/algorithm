package main

type Weight interface {
	Less(w Weight) bool
	Add(w Weight) Weight
}

type IntWeight int

func (iw IntWeight) Less(iw1 Weight) bool {
	val := iw1.(IntWeight)
	return iw < val
}

func (iw IntWeight) Add(iw1 Weight) Weight {
	val := iw1.(IntWeight)
	return iw + val
}

// 代表a->b的边
type edge struct {
	a, b   int
	weight Weight
}

func newEdge(a, b int, w Weight) *edge {
	var e = edge{
		a:      a,
		b:      b,
		weight: w,
	}
	return &e
}

func (e *edge) v() int {
	return e.a
}

func (e *edge) w() int {
	return e.b
}

func (e *edge) wt() Weight {
	return e.weight
}

func (e *edge) less(e1 *edge) bool {
	return e.weight.Less(e1.weight)
}

func (e *edge) other(x int) int {
	if x != e.a && x != e.b {
		panic("invalid node")
	}

	if e.a == x {
		return e.b
	}
	return e.a
}
