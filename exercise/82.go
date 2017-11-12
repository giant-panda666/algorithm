package main

//Given a sorted linked list, delete all nodes that have duplicate numbers, leaving only distinct numbers from the original list.
//
//For example,
//Given 1->2->3->3->4->4->5, return 1->2->5.
//Given 1->1->1->2->3, return 2->3.

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var vhead = new(ListNode)
	vhead.Next = head
	var p0, p1, p2 = vhead, head, head.Next
	for p2 != nil {
		if p1.Val == p2.Val {
			for p2 != nil && p1.Val == p2.Val {
				p2 = p2.Next
			}
			p0.Next = p2
			p1 = p2
			if p2 == nil {
				return vhead.Next
			}
			p2 = p2.Next
		} else {
			p0 = p1
			p1 = p2
			p2 = p2.Next
		}
	}
	return vhead.Next
}
