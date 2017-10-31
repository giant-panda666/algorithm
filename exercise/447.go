package main

import "fmt"

//Given n points in the plane that are all pairwise distinct, a "boomerang" is a tuple of points (i, j, k) such that the distance between i and j equals the distance between i and k (the order of the tuple matters).
//
//Find the number of boomerangs. You may assume that n will be at most 500 and coordinates of points are all in the range [-10000, 10000] (inclusive).
//
//Example:
//Input:
//[[0,0],[1,0],[2,0]]
//
//Output:
//2
//
//Explanation:
//The two boomerangs are [[1,0],[0,0],[2,0]] and [[1,0],[2,0],[0,0]]

func numberOfBoomerangs(points [][]int) int {
	var res int
	for i := 0; i < len(points); i++ {
		var record = make(map[int]int)
		for j := 0; j < len(points); j++ {
			if j != i {
				dis := (points[i][0]-points[j][0])*(points[i][0]-points[j][0]) + (points[i][1]-points[j][1])*(points[i][1]-points[j][1])
				record[dis]++
			}
		}
		for _, v := range record {
			res += v * (v - 1)
		}
	}

	return res
}

func main() {
	var points = [][]int{[]int{0, 0}, []int{1, 0}, []int{2, 0}}
	fmt.Println("expect 2, result:", numberOfBoomerangs(points))
}
