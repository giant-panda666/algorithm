package main

//Given a binary tree containing digits from 0-9 only, each root-to-leaf path could represent a number.
//
//An example is the root-to-leaf path 1->2->3 which represents the number 123.
//
//Find the total sum of all root-to-leaf numbers.
//
//For example,
//
//    1
//   / \
//  2   3
//The root-to-leaf path 1->2 represents the number 12.
//The root-to-leaf path 1->3 represents the number 13.
//
//Return the sum = 12 + 13 = 25.
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
	var res = findAllNumbers(root)
	var sum int
	for _, v := range res {
		vv, _ := strconv.Atoi(v)
		sum += vv
	}
	return sum
}

func findAllNumbers(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return []string{fmt.Sprintf("%d", root.Val)}
	}

	var res []string
	for _, v := range findAllNumbers(root.Left) {
		tmp := fmt.Sprintf("%d%s", root.Val, v)
		res = append(res, tmp)
	}
	for _, v := range findAllNumbers(root.Right) {
		tmp := fmt.Sprintf("%d%s", root.Val, v)
		res = append(res, tmp)
	}

	return res
}
