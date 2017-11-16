package main

//Given a binary tree, determine if it is height-balanced.
//
//For this problem, a height-balanced binary tree is defined as a binary tree in which the depth of the two subtrees of every node never differ by more than 1.
//

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	lhi := maxHeight(root.Left)
	rhi := maxHeight(root.Right)
	if lhi-rhi > 1 || rhi-lhi > 1 {
		return false
	}

	return isBalanced(root.Left) && isBalanced(root.Right)
}

func maxHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max(1+maxHeight(root.Left), 1+maxHeight(root.Right))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
