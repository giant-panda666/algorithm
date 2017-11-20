package main

import (
	"fmt"
	"sort"
)

//Given a collection of candidate numbers (C) and a target number (T), find all unique combinations in C where the candidate numbers sums to T.
//
//Each number in C may only be used once in the combination.
//
//Note:
//All numbers (including target) will be positive integers.
//The solution set must not contain duplicate combinations.
//For example, given candidate set [10, 1, 2, 7, 6, 1, 5] and target 8,
//A solution set is:
//[
//  [1, 7],
//  [1, 2, 5],
//  [2, 6],
//  [1, 1, 6]
//]
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var c cs2
	c.generateCombinationSum2(candidates, 0, target)
	return c.res
}

type cs2 struct {
	cur    []int
	curSum int
	res    [][]int
}

func (c *cs2) generateCombinationSum2(candidates []int, index, target int) {
	if c.curSum == target {
		tmp := make([]int, len(c.cur))
		copy(tmp, c.cur)
		c.res = append(c.res, tmp)
		return
	}
	for i := index; i < len(candidates); i++ {
		if i > index && candidates[i] == candidates[i-1] {
			continue
		}
		v := candidates[i]
		if c.curSum+v > target {
			return
		}
		c.cur = append(c.cur, v)
		c.curSum += v
		c.generateCombinationSum2(candidates, i+1, target)
		c.cur = c.cur[:len(c.cur)-1]
		c.curSum -= v
	}
}

func main() {
	var candidates = []int{10, 1, 2, 7, 6, 1, 5}
	var target = 8
	fmt.Println(combinationSum2(candidates, target))
}
