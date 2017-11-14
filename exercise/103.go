package main

//Given a binary tree, return the zigzag level order traversal of its nodes' values. (ie, from left to right, then right to left for the next level and alternate between).
//
//For example:
//Given binary tree [3,9,20,null,null,15,7],
//
//    3
//   / \
//  9  20
//    /  \
//   15   7
//return its zigzag level order traversal as:
//[
//  [3],
//  [20,9],
//  [15,7]
//]

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type pair struct {
	node  *TreeNode
	level int
}

type queue struct {
	data  []pair
	count int
}

func (q *queue) push(p pair) {
	q.data = append(q.data, p)
	q.count++
}

func (q *queue) pop() pair {
	q.count--
	res := q.data[0]
	q.data = q.data[1:]
	return res
}

func (q *queue) isEmpty() bool {
	return q.count == 0
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var res [][]int
	var q queue
	q.push(pair{node: root, level: 0})
	for !q.isEmpty() {
		tmp := q.pop()
		level := tmp.level
		node := tmp.node
		if len(res) <= level {
			res = append(res, []int{node.Val})
		} else {
			if level%2 == 0 {
				res[level] = append(res[level], node.Val)
			} else {
				res[level] = append([]int{node.Val}, res[level][:]...)
			}
		}

		if node.Left != nil {
			q.push(pair{node: node.Left, level: level + 1})
		}
		if node.Right != nil {
			q.push(pair{node: node.Right, level: level + 1})
		}
	}

	return res
}
