package main

//Given a binary tree, determine if it is a valid binary search tree (BST).
//
//Assume a BST is defined as follows:
//
//The left subtree of a node contains only nodes with keys less than the node's key.
//The right subtree of a node contains only nodes with keys greater than the node's key.
//Both the left and right subtrees must also be binary search trees.
//Example 1:
//    2
//   / \
//  1   3
//Binary tree [2,1,3], return true.
//Example 2:
//    1
//   / \
//  2   3
//Binary tree [1,2,3], return false.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if root.Left != nil {
		if !isValidBST(root.Left) || !isLessThan(root.Left, root.Val) {
			return false
		}
	}
	if root.Right != nil {
		if !isValidBST(root.Right) || !isMoreThan(root.Right, root.Val) {
			return false
		}
	}

	return true
}

func isLessThan(root *TreeNode, val int) bool {
	if root == nil {
		return true
	}
	if root.Val >= val {
		return false
	}
	return isLessThan(root.Left, val) && isLessThan(root.Right, val)
}

func isMoreThan(root *TreeNode, val int) bool {
	if root == nil {
		return true
	}
	if root.Val <= val {
		return false
	}
	return isMoreThan(root.Right, val) && isMoreThan(root.Left, val)
}
