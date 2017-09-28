package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func generateTestArr(n, l, r int) []int {
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

func generateNearlyOrderedTestArr(n, l, r int) []int {
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

func printArr(a []int) {
	for _, v := range a {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

func TestSelectionSort(t *testing.T) {
	s1 := generateTestArr(10, 1, 10)
	benchmark("selectionSort", selectionSort, s1)
	printArr(s1)

	s2 := generateTestArr(100, 1, 100)
	benchmark("selectionSort", selectionSort, s2)

	s3 := generateTestArr(1000, 1, 1000)
	benchmark("selectionSort", selectionSort, s3)

	s4 := generateTestArr(10000, 1, 10000)
	benchmark("selectionSort", selectionSort, s4)
}

func TestInsertionSort(t *testing.T) {
	s1 := generateTestArr(10, 1, 10)
	benchmark("insertSort", insertSort, s1)

	s2 := generateTestArr(100, 1, 100)
	benchmark("insertSort", insertSort, s2)

	s3 := generateTestArr(1000, 1, 1000)
	benchmark("insertSort", insertSort, s3)

	s4 := generateTestArr(10000, 1, 10000)
	benchmark("insertSort", insertSort, s4)
}

func TestInsertionSortAdvance(t *testing.T) {
	s1 := generateTestArr(10, 1, 10)
	benchmark("insertSortAdvance", insertSortAdvance, s1)
	printArr(s1)

	s2 := generateTestArr(100, 1, 100)
	benchmark("insertSortAdvance", insertSortAdvance, s2)

	s3 := generateTestArr(1000, 1, 1000)
	benchmark("insertSortAdvance", insertSortAdvance, s3)

	s4 := generateTestArr(10000, 1, 10000)
	benchmark("insertSortAdvance", insertSortAdvance, s4)
	s44 := generateNearlyOrderedTestArr(10000, 1, 10000)
	benchmark("insertSortAdvance", insertSortAdvance, s44)
}

func TestBobbleSort(t *testing.T) {
	s1 := generateTestArr(10, 1, 10)
	benchmark("bobbleSort", bobbleSort, s1)
	printArr(s1)

	s2 := generateTestArr(100, 1, 100)
	benchmark("bobbleSort", bobbleSort, s2)

	s3 := generateTestArr(1000, 1, 1000)
	benchmark("bobbleSort", bobbleSort, s3)

	s4 := generateTestArr(10000, 1, 10000)
	benchmark("bobbleSort", bobbleSort, s4)
	s44 := generateNearlyOrderedTestArr(10000, 1, 10000)
	benchmark("bobbleSort", bobbleSort, s44)
}

func TestBobbleSortAdvance(t *testing.T) {
	s1 := generateTestArr(10, 1, 10)
	benchmark("bobbleSortAdvance", bobbleSortAdvance, s1)
	printArr(s1)

	s2 := generateTestArr(100, 1, 100)
	benchmark("bobbleSortAdvance", bobbleSortAdvance, s2)

	s3 := generateTestArr(1000, 1, 1000)
	benchmark("bobbleSortAdvance", bobbleSortAdvance, s3)

	s4 := generateTestArr(10000, 1, 10000)
	benchmark("bobbleSortAdvance", bobbleSortAdvance, s4)
	s44 := generateNearlyOrderedTestArr(10000, 1, 10000)
	benchmark("bobbleSortAdvance", bobbleSortAdvance, s44)
}

func TestShellSort(t *testing.T) {
	s1 := generateTestArr(10, 1, 10)
	benchmark("shellSort", shellSort, s1)
	printArr(s1)

	s2 := generateTestArr(100, 1, 100)
	benchmark("shellSort", shellSort, s2)

	s3 := generateTestArr(1000, 1, 1000)
	benchmark("shellSort", shellSort, s3)

	s4 := generateTestArr(10000, 1, 10000)
	benchmark("shellSort", shellSort, s4)
	s44 := generateNearlyOrderedTestArr(10000, 1, 10000)
	benchmark("shellSort", shellSort, s44)
}

//func TestSelectionSortUsingTemplate(t *testing.T) {
//	s := generateTestArr(10, 1, 10)
//	var a intSlice = s
//	fmt.Println(selectionSortUsingTemplate(a))
//}

func benchmark(name string, f func([]int) []int, a []int) {
	start := time.Now()
	f(a)
	fmt.Println(name+" Using: ", time.Since(start))
}
