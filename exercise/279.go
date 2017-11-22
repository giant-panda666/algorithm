package main

import (
	"fmt"
	"math"
)

//Given a positive integer n, find the least number of perfect square numbers (for example, 1, 4, 9, 16, ...) which sum to n.
//
//For example, given n = 12, return 3 because 12 = 4 + 4 + 4; given n = 13, return 2 because 13 = 4 + 9.

type pair struct {
	first, second int
}

type queue struct {
	data  []*pair
	count int
}

func (q *queue) push(p *pair) {
	q.data = append(q.data, p)
	q.count++
}

func (q *queue) pop() *pair {
	if q.count == 0 {
		return nil
	}
	q.count--
	res := q.data[0]
	q.data = q.data[1:]
	return res
}

func (q *queue) isEmpty() bool {
	return q.count == 0
}

func numSquares(n int) int {
	var q queue
	q.push(&pair{n, 0})

	var visited = make([]bool, n+1)
	visited[n] = true

	for !q.isEmpty() {
		tmp := q.pop()
		num := tmp.first
		step := tmp.second

		if num == 0 {
			return step
		}

		for i := 1; num-i*i >= 0; i++ {
			if !visited[num-i*i] {
				q.push(&pair{num - i*i, step + 1})
				visited[num-i*i] = true
			}
		}
	}

	return 0
}

func numSquares2(n int) int {
	if n == 1 {
		return 1
	}

	var memo = make([]int, n+1)
	for i := 1; i <= n; i++ {
		memo[i] = i
		for j := 1; j*j <= i; j++ {
			if memo[i] > memo[i-j*j]+1 {
				memo[i] = memo[i-j*j] + 1
			}
		}
	}
	fmt.Println(memo)

	return memo[n]
}

func main() {
	fmt.Println("numSquares1(12)", numSquares(12))
	fmt.Println("numSquares1(13)", numSquares(13))
	fmt.Println("numSquares2(12)", numSquares2(12))
	fmt.Println("numSquares2(13)", numSquares2(13))
}
