package main

//Given a binary tree, return the postorder traversal of its nodes' values.
//
//For example:
//Given binary tree {1,#,2,3},
//   1
//    \
//     2
//    /
//   3
//return [3,2,1].

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal1(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var res []int
	res = append(res, postorderTraversal1(root.Left)...)
	res = append(res, postorderTraversal1(root.Right)...)
	res = append(res, root.Val)
	return res
}

type Command struct {
	command string // "go" "print"
	node    *TreeNode
}

func newCommand(command string, node *TreeNode) *Command {
	return &Command{
		command: command,
		node:    node,
	}
}

type stack struct {
	data  []*Command
	count int
}

func (s *stack) push(c *Command) {
	s.data = append(s.data, c)
	s.count++
}

func (s *stack) pop() *Command {
	if s.count == 0 {
		return nil
	}
	s.count--
	res := s.data[s.count]
	s.data = s.data[:s.count]
	return res
}

func (s *stack) isEmpty() bool {
	return s.count == 0
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var s stack
	var c = newCommand("go", root)
	s.push(c)

	var res []int
	for !s.isEmpty() {
		tmp := s.pop()
		if tmp.command == "print" {
			res = append(res, tmp.node.Val)
		} else {
			s.push(newCommand("print", tmp.node))
			if tmp.node.Right != nil {
				s.push(newCommand("go", tmp.node.Right))
			}
			if tmp.node.Left != nil {
				s.push(newCommand("go", tmp.node.Left))
			}
		}
	}

	return res
}
