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

	s3 := util.GenerateTestArr(1000, 1, 1000)
	util.Benchmark("mergeSort", mergeSort, s3)

	s4 := util.GenerateTestArr(10000, 1, 10000)
	util.Benchmark("mergeSort", mergeSort, s4)
	s44 := util.GenerateNearlyOrderedTestArr(10000, 1, 10000)
	util.Benchmark("mergeSort", mergeSort, s44)
}

func TestMergeSortBottomup(t *testing.T) {
	s1 := util.GenerateTestArr(10, 1, 10)
	util.Benchmark("mergeSortBottomup", mergeSortBottomup, s1)
	util.PrintArr(s1)

	s2 := util.GenerateTestArr(100, 1, 100)
	util.Benchmark("mergeSortBottomup", mergeSortBottomup, s2)

	s3 := util.GenerateTestArr(1000, 1, 1000)
	util.Benchmark("mergeSortBottomup", mergeSortBottomup, s3)

	s4 := util.GenerateTestArr(10000, 1, 10000)
	util.Benchmark("mergeSortBottomup", mergeSortBottomup, s4)
	s44 := util.GenerateNearlyOrderedTestArr(10000, 1, 10000)
	util.Benchmark("mergeSortBottomup", mergeSortBottomup, s44)
}
