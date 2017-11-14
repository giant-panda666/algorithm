package main

//Given a binary tree, return the bottom-up level order traversal of its nodes' values. (ie, from left to right, level by level from leaf to root).
//
//For example:
//Given binary tree [3,9,20,null,null,15,7],
//    3
//   / \
//  9  20
//    /  \
//   15   7
//return its bottom-up level order traversal as:
//[
//  [15,7],
//  [9,20],
//  [3]
//]

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var res [][]int
	var treeNodes = []*TreeNode{root}
	for len(treeNodes) > 0 {
		var childTreeNodes []*TreeNode
		var values []int
		for i := 0; i < len(treeNodes); i++ {
			values = append(values, treeNodes[i].Val)
			if treeNodes[i].Left != nil {
				childTreeNodes = append(childTreeNodes, treeNodes[i].Left)
			}
			if treeNodes[i].Right != nil {
				childTreeNodes = append(childTreeNodes, treeNodes[i].Right)
			}
		}
		res = append(res, values)
		treeNodes = childTreeNodes
	}

	for i, j := len(res)-1, 0; i > j; i, j = i-1, j+1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}
