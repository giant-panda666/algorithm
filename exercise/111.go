package main

//Given a binary tree, find its minimum depth.
//
//The minimum depth is the number of nodes along the shortest path from the root node down to the nearest leaf node.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}
	var minL, minR int
	if root.Left != nil {
		minL = minDepth(root.Left) + 1
	}
	if root.Right != nil {
		minR = minDepth(root.Right) + 1
	}
	if minL == 0 {
		return minR
	}
	if minR == 0 {
		return minL
	}
	return min(minL, minR)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
