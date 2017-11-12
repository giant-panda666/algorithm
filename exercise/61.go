package main

//Given a list, rotate the list to the right by k places, where k is non-negative.

//Example:
//
//Given 1->2->3->4->5->NULL and k = 2,
//
//return 4->5->1->2->3->NULL.

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if k <= 0 || head == nil {
		return head
	}
	var len int
	for l := head; l != nil; l = l.Next {
		len++
	}
	k = k % len

	var p, q = head, head
	for i := 0; i < k; i++ {
		q = q.Next
	}
	for q.Next != nil {
		q = q.Next
		p = p.Next
	}

	q.Next = head
	head = p.Next
	p.Next = nil
	return head
}
