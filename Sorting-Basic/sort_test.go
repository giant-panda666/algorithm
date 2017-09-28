package main

import (
	"algorithm/util"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	s1 := util.GenerateTestArr(10, 1, 10)
	util.Benchmark("selectionSort", selectionSort, s1)
	util.PrintArr(s1)

	s2 := util.GenerateTestArr(100, 1, 100)
	util.Benchmark("selectionSort", selectionSort, s2)

	s3 := util.GenerateTestArr(1000, 1, 1000)
	util.Benchmark("selectionSort", selectionSort, s3)

	s4 := util.GenerateTestArr(10000, 1, 10000)
	util.Benchmark("selectionSort", selectionSort, s4)
}

func TestInsertionSort(t *testing.T) {
	s1 := util.GenerateTestArr(10, 1, 10)
	util.Benchmark("insertSort", insertSort, s1)

	s2 := util.GenerateTestArr(100, 1, 100)
	util.Benchmark("insertSort", insertSort, s2)

	s3 := util.GenerateTestArr(1000, 1, 1000)
	util.Benchmark("insertSort", insertSort, s3)

	s4 := util.GenerateTestArr(10000, 1, 10000)
	util.Benchmark("insertSort", insertSort, s4)
}

func TestInsertionSortAdvance(t *testing.T) {
	s1 := util.GenerateTestArr(10, 1, 10)
	util.Benchmark("insertSortAdvance", insertSortAdvance, s1)
	util.PrintArr(s1)

	s2 := util.GenerateTestArr(100, 1, 100)
	util.Benchmark("insertSortAdvance", insertSortAdvance, s2)

	s3 := util.GenerateTestArr(1000, 1, 1000)
	util.Benchmark("insertSortAdvance", insertSortAdvance, s3)

	s4 := util.GenerateTestArr(10000, 1, 10000)
	util.Benchmark("insertSortAdvance", insertSortAdvance, s4)
	s44 := util.GenerateNearlyOrderedTestArr(10000, 1, 10000)
	util.Benchmark("insertSortAdvance", insertSortAdvance, s44)
}

func TestBobbleSort(t *testing.T) {
	s1 := util.GenerateTestArr(10, 1, 10)
	util.Benchmark("bobbleSort", bobbleSort, s1)
	util.PrintArr(s1)

	s2 := util.GenerateTestArr(100, 1, 100)
	util.Benchmark("bobbleSort", bobbleSort, s2)

	s3 := util.GenerateTestArr(1000, 1, 1000)
	util.Benchmark("bobbleSort", bobbleSort, s3)

	s4 := util.GenerateTestArr(10000, 1, 10000)
	util.Benchmark("bobbleSort", bobbleSort, s4)
	s44 := util.GenerateNearlyOrderedTestArr(10000, 1, 10000)
	util.Benchmark("bobbleSort", bobbleSort, s44)
}

func TestBobbleSortAdvance(t *testing.T) {
	s1 := util.GenerateTestArr(10, 1, 10)
	util.Benchmark("bobbleSortAdvance", bobbleSortAdvance, s1)
	util.PrintArr(s1)

	s2 := util.GenerateTestArr(100, 1, 100)
	util.Benchmark("bobbleSortAdvance", bobbleSortAdvance, s2)

	s3 := util.GenerateTestArr(1000, 1, 1000)
	util.Benchmark("bobbleSortAdvance", bobbleSortAdvance, s3)

	s4 := util.GenerateTestArr(10000, 1, 10000)
	util.Benchmark("bobbleSortAdvance", bobbleSortAdvance, s4)
	s44 := util.GenerateNearlyOrderedTestArr(10000, 1, 10000)
	util.Benchmark("bobbleSortAdvance", bobbleSortAdvance, s44)
}

func TestShellSort(t *testing.T) {
	s1 := util.GenerateTestArr(10, 1, 10)
	util.Benchmark("shellSort", shellSort, s1)
	util.PrintArr(s1)

	s2 := util.GenerateTestArr(100, 1, 100)
	util.Benchmark("shellSort", shellSort, s2)

	s3 := util.GenerateTestArr(1000, 1, 1000)
	util.Benchmark("shellSort", shellSort, s3)

	s4 := util.GenerateTestArr(10000, 1, 10000)
	util.Benchmark("shellSort", shellSort, s4)
	s44 := util.GenerateNearlyOrderedTestArr(10000, 1, 10000)
	util.Benchmark("shellSort", shellSort, s44)
}

//func TestSelectionSortUsingTemplate(t *testing.T) {
//	s := util.GenerateTestArr(10, 1, 10)
//	var a intSlice = s
//	fmt.Println(selectionSortUsingTemplate(a))
//}
