package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestUnionFind(t *testing.T) {
	testN(10000)
	//testN(100000)
}

func testN(n int) {
	var u = newUnionFind(n)
	var qu = newQuickUnionFind(n)
	rand.Seed(time.Now().Unix())

	s := time.Now()
	for i := 0; i < n; i++ {
		a := rand.Int() % n
		b := rand.Int() % n
		u.union(a, b)
	}
	for i := 0; i < n; i++ {
		a := rand.Int() % n
		b := rand.Int() % n
		u.isConnected(a, b)
	}
	fmt.Println("unionFind, n:", n, "using time:", time.Since(s))

	s = time.Now()
	for i := 0; i < n; i++ {
		a := rand.Int() % n
		b := rand.Int() % n
		qu.union(a, b)
	}
	for i := 0; i < n; i++ {
		a := rand.Int() % n
		b := rand.Int() % n
		qu.isConnected(a, b)
	}
	fmt.Println("quickUnionFind, n:", n, "using time:", time.Since(s))
}
