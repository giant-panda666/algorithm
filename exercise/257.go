package main

//Given a binary tree, return all root-to-leaf paths.
//
//For example, given the following binary tree:
//
//   1
// /   \
//2     3
// \
//  5
//All root-to-leaf paths are:
//
//["1->2->5", "1->3"]

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}

	var res []string
	if root.Left == nil && root.Right == nil {
		return []string{strconv.Itoa(root.Val)}
	}
	if root.Left != nil {
		for _, v := range binaryTreePaths(root.Left) {
			res = append(res, fmt.Sprintf("%s->%s", strconv.Itoa(root.Val), v))
		}
	}
	if root.Right != nil {
		rightRes := binaryTreePaths(root.Right)
		for _, v := range binaryTreePaths(root.Right) {
			res = append(res, fmt.Sprintf("%s->%s", strconv.Itoa(root.Val), v))
		}
	}
	return res
}
