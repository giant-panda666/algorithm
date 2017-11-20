package main

import (
	"fmt"
	"sort"
)

//Given a set of candidate numbers (C) (without duplicates) and a target number (T), find all unique combinations in C where the candidate numbers sums to T.
//
//The same repeated number may be chosen from C unlimited number of times.
//
//Note:
//All numbers (including target) will be positive integers.
//The solution set must not contain duplicate combinations.
//For example, given candidate set [2, 3, 6, 7] and target 7,
//A solution set is:
//[
//  [7],
//  [2, 2, 3]
//]

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var c cs
	c.generateCombinationSum(candidates, 0, target)
	return c.res
}

type cs struct {
	cur    []int
	curSum int
	res    [][]int
}

func (c *cs) generateCombinationSum(candidates []int, index int, target int) {
	if c.curSum == target {
		tmp := make([]int, len(c.cur))
		copy(tmp, c.cur)
		c.res = append(c.res, tmp)
		return
	}
	for i := index; i < len(candidates); i++ {
		v := candidates[i]
		if c.curSum+v > target {
			return
		}
		c.cur = append(c.cur, v)
		c.curSum += v
		c.generateCombinationSum(candidates, i, target)
		c.cur = c.cur[:len(c.cur)-1]
		c.curSum -= v
	}
}

func main() {
	var candidates = []int{2, 3, 6, 7}
	var target = 7
	fmt.Println(combinationSum(candidates, target))
}
