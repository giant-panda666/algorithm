package main

//Given a binary tree and a sum, find all root-to-leaf paths where each path's sum equals the given sum.
//
//For example:
//Given the below binary tree and sum = 22,
//              5
//             / \
//            4   8
//           /   / \
//          11  13  4
//         /  \    / \
//        7    2  5   1
//return
//[
//   [5,4,11,2],
//   [5,8,4,5]
//]

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil && sum == root.Val {
		return [][]int{[]int{root.Val}}
	}

	var res [][]int
	for _, v := range pathSum(root.Left, sum-root.Val) {
		v = append([]int{root.Val}, v...)
		res = append(res, v)
	}
	for _, v := range pathSum(root.Right, sum-root.Val) {
		v = append([]int{root.Val}, v...)
		res = append(res, v)
	}

	return res
}
