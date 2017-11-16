package main

//Given a binary search tree, write a function kthSmallest to find the kth smallest element in it.
//
//Note:
//You may assume k is always valid, 1 ≤ k ≤ BST's total elements.
//
//Follow up:
//What if the BST is modified (insert/delete operations) often and you need to find the kth smallest frequently? How would you optimize the kthSmallest routine?

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
	var order []int
	findOrder(root, &order)
	return order[k-1]
}

func findOrder(root *TreeNode, order *[]int) {
	if root == nil {
		return
	}
	findOrder(root.Left, order)
	*order = append(*order, root.Val)
	findOrder(root.Right, order)
}
