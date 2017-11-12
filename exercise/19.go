package main

//Given a linked list, remove the nth node from the end of list and return its head.
//
//For example,
//
//   Given linked list: 1->2->3->4->5, and n = 2.
//
//   After removing the second node from the end, the linked list becomes 1->2->3->5.
//Note:
//Given n will always be valid.
//Try to do this in one pass.

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var len int
	for l := head; l != nil; l = l.Next {
		len++
	}

	var vhead = new(ListNode)
	vhead.Next = head
	var prev, cur *ListNode = vhead, head
	for i := 0; i < len-n; i++ {
		prev = cur
		cur = cur.Next
	}

	prev.Next = cur.Next
	cur.Next = nil
	return vhead.Next
}

func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	if n < 0 {
		return head
	}

	var vhead = new(ListNode)
	vhead.Next = head

	var p, q = vhead, vhead
	for i := 0; i < n+1; i++ {
		q = q.Next
	}
	for q != nil {
		q = q.Next
		p = p.Next
	}

	p.Next = p.Next.Next
	return vhead.Next
}
