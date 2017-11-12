package main

//Given a linked list, swap every two adjacent nodes and return its head.
//
//For example,
//Given 1->2->3->4, you should return the list as 2->1->4->3.
//
//Your algorithm should use only constant space. You may not modify the values in the list, only nodes itself can be changed.

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var vhead = new(ListNode)
	var p, n1, n2 = vhead, head, head.Next
	for n2 != nil {
		p.Next = n2
		n1.Next = n2.Next
		n2.Next = n1

		p = n1
		n1 = n1.Next
		if n1 == nil {
			return vhead.Next
		}
		n2 = n1.Next
	}
	return vhead.Next
}
