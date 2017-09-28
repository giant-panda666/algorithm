package util

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateTestArr(n, l, r int) []int {
	if n <= 0 || l > r {
		return nil
	}

	rand.Seed(time.Now().Unix())

	var ret []int
	for i := 0; i < n; i++ {
		tmp := rand.Int()%(r-l+1) + l
		ret = append(ret, tmp)
	}

	return ret
}

func GenerateNearlyOrderedTestArr(n, l, r int) []int {
	var ret []int
	for i := 0; i < n; i++ {
		ret = append(ret, i+l)
	}

	var c = rand.Int() % 100
	for i := 0; i < c; i++ {
		x := rand.Int() % n
		y := rand.Int() % n
		ret[x], ret[y] = ret[y], ret[x]
	}

	return ret
}

func PrintArr(a []int) {
	for _, v := range a {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

func Benchmark(name string, f func([]int) []int, a []int) {
	start := time.Now()
	f(a)
	fmt.Println(name+" Using: ", time.Since(start))
}
