package main

//Given a binary tree, imagine yourself standing on the right side of it, return the values of the nodes you can see ordered from top to bottom.
//
//For example:
//Given the following binary tree,
//   1            <---
// /   \
//2     3         <---
// \     \
//  5     4       <---
//You should return [1, 3, 4].
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var res []int
	var treeNodes = []*TreeNode{root}
	for len(treeNodes) > 0 {
		childTreeNodes := []*TreeNode{}
		count := -1
		for i := 0; i < len(treeNodes); i++ {
			if treeNodes[i].Left != nil {
				childTreeNodes = append(childTreeNodes, treeNodes[i].Left)
			}
			if treeNodes[i].Right != nil {
				childTreeNodes = append(childTreeNodes, treeNodes[i].Right)
			}
			count = i
		}
		if count > -1 {
			res = append(res, treeNodes[count].Val)
		} else {
			return res
		}
		treeNodes = childTreeNodes
	}

	return res
}
