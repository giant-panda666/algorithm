package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestUnionFind(t *testing.T) {
	testN1(10000)
	testN2(100000)
	testN3(100000)
	testN4(100000)
}

func testN1(n int) {
	var u = newUnionFind(n)
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
}

func testN2(n int) {
	var qu = newQuickUnionFindSize(n)
	rand.Seed(time.Now().Unix())

	s := time.Now()
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
	fmt.Println("quickUnionFindSize, n:", n, "using time:", time.Since(s))
}

func testN3(n int) {
	var qu = newQuickUnionFindRank(n)
	rand.Seed(time.Now().Unix())

	s := time.Now()
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
	fmt.Println("quickUnionFindRank, n:", n, "using time:", time.Since(s))
}

func testN4(n int) {
	var qu = newQuickUnionFindPathCompression(n)
	rand.Seed(time.Now().Unix())

	s := time.Now()
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
	fmt.Println("quickUnionFindPathCompression, n:", n, "using time:", time.Since(s))
}
