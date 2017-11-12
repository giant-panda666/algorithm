package main

//You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.
//
//You may assume the two numbers do not contain any leading zero, except the number 0 itself.
//
//Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
//Output: 7 -> 0 -> 8
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var n = new(ListNode)
	var p = n
	var nextVal = 0
	for l1 != nil && l2 != nil {
		val := l1.Val + l2.Val + nextVal
		tmp := new(ListNode)
		tmp.Val = val % 10
		nextVal = val / 10
		p.Next = tmp
		p = tmp
		l1 = l1.Next
		l2 = l2.Next
	}
	if l1 != nil {
		for l1 != nil {
			val := l1.Val + nextVal
			tmp := new(ListNode)
			tmp.Val = val % 10
			nextVal = val / 10
			p.Next = tmp
			p = tmp
			l1 = l1.Next
		}
	}
	if l2 != nil {
		for l2 != nil {
			val := l2.Val + nextVal
			tmp := new(ListNode)
			tmp.Val = val % 10
			nextVal = val / 10
			p.Next = tmp
			p = tmp
			l2 = l2.Next
		}
	}

	if nextVal != 0 {
		p.Next = &ListNode{
			Val: nextVal,
		}
	}

	return n.Next
}
