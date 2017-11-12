package main

//Given a linked list, reverse the nodes of a linked list k at a time and return its modified list.
//
//k is a positive integer and is less than or equal to the length of the linked list. If the number of nodes is not a multiple of k then left-out nodes in the end should remain as it is.
//
//You may not alter the values in the nodes, only nodes itself may be changed.
//
//Only constant memory is allowed.
//
//For example,
//Given this linked list: 1->2->3->4->5
//
//For k = 2, you should return: 2->1->4->3->5
//
//For k = 3, you should return: 3->2->1->4->5

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	var dummyHead = new(ListNode)
	dummyHead.Next = head
	var prev, start, end *ListNode = dummyHead, head, head
	for end != nil {
		for i := 0; i < k; i++ {
			if end == nil {
				return dummyHead.Next
			} else {
				end = end.Next
			}
		}
		prev.Next = nil

		var cur, last *ListNode
		for start != end {
			prev.Next = start
			tmp := start.Next
			start.Next = cur
			if cur == nil {
				last = start
			}
			cur = start
			start = tmp
		}
		last.Next = start
		prev = last
	}
	return dummyHead.Next
}
