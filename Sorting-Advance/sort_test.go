package main

import (
	"algorithm/util"
	"testing"
)

func TestMergeSort(t *testing.T) {
	s1 := util.GenerateTestArr(10, 1, 10)
	util.Benchmark("mergeSort", mergeSort, s1)
	util.PrintArr(s1)

	s2 := util.GenerateTestArr(100, 1, 100)
	util.Benchmark("mergeSort", mergeSort, s2)
	util.IsOrdered(s2)

	s3 := util.GenerateTestArr(1000, 1, 1000)
	util.Benchmark("mergeSort", mergeSort, s3)
	util.IsOrdered(s3)

	s4 := util.GenerateTestArr(10000, 1, 10000)
	util.Benchmark("mergeSort", mergeSort, s4)
	util.IsOrdered(s4)
	s44 := util.GenerateNearlyOrderedTestArr(10000, 1, 10000)
	util.Benchmark("mergeSort", mergeSort, s44)
	util.IsOrdered(s44)
}

func TestMergeSortBottomup(t *testing.T) {
	s1 := util.GenerateTestArr(10, 1, 10)
	util.Benchmark("mergeSortBottomup", mergeSortBottomup, s1)
	util.PrintArr(s1)

	s2 := util.GenerateTestArr(100, 1, 100)
	util.Benchmark("mergeSortBottomup", mergeSortBottomup, s2)
	util.IsOrdered(s2)

	s3 := util.GenerateTestArr(1000, 1, 1000)
	util.Benchmark("mergeSortBottomup", mergeSortBottomup, s3)
	util.IsOrdered(s3)

	s4 := util.GenerateTestArr(10000, 1, 10000)
	util.Benchmark("mergeSortBottomup", mergeSortBottomup, s4)
	util.IsOrdered(s4)
	s44 := util.GenerateNearlyOrderedTestArr(10000, 1, 10000)
	util.Benchmark("mergeSortBottomup", mergeSortBottomup, s44)
	util.IsOrdered(s44)
}

func TestQuickSort(t *testing.T) {
	s1 := util.GenerateTestArr(10, 1, 10)
	util.Benchmark("quickSort", quickSort, s1)
	util.PrintArr(s1)

	s2 := util.GenerateTestArr(100, 1, 100)
	util.Benchmark("quickSort", quickSort, s2)
	util.IsOrdered(s2)

	s3 := util.GenerateTestArr(1000, 1, 1000)
	util.Benchmark("quickSort", quickSort, s3)
	util.IsOrdered(s3)

	s4 := util.GenerateTestArr(10000, 1, 10000)
	util.Benchmark("quickSort", quickSort, s4)
	util.IsOrdered(s4)
	s44 := util.GenerateNearlyOrderedTestArr(10000, 1, 10000)
	util.Benchmark("quickSort", quickSort, s44)
	util.IsOrdered(s44)

	//	s45 := util.GenerateNearlyOrderedTestArr(1000000, 1, 1000000)
	//	util.Benchmark("quickSort", quickSort, s45)
}

func TestQuickSortRandom(t *testing.T) {
	s1 := util.GenerateTestArr(10, 1, 10)
	util.Benchmark("quickSortRandom", quickSortRandom, s1)
	util.PrintArr(s1)

	s2 := util.GenerateTestArr(100, 1, 100)
	util.Benchmark("quickSortRandom", quickSortRandom, s2)
	util.IsOrdered(s2)

	s3 := util.GenerateTestArr(1000, 1, 1000)
	util.Benchmark("quickSortRandom", quickSortRandom, s3)
	util.IsOrdered(s3)

	s4 := util.GenerateTestArr(10000, 1, 10000)
	util.Benchmark("quickSortRandom", quickSortRandom, s4)
	util.IsOrdered(s4)
	s44 := util.GenerateNearlyOrderedTestArr(10000, 1, 10000)
	util.Benchmark("quickSortRandom", quickSortRandom, s44)
	util.IsOrdered(s44)

	s45 := util.GenerateNearlyOrderedTestArr(1000000, 1, 1000000)
	util.Benchmark("quickSortRandom", quickSortRandom, s45)
	util.IsOrdered(s45)
}
