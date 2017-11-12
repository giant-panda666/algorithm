package main

//Sort a linked list using insertion sort.

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var vhead = new(ListNode)
	vhead.Next = head
	todo := head.Next
	head.Next = nil

	for todo != nil {
		tmp := todo
		todo = todo.Next

		prev, cur := vhead, vhead.Next
		for cur != nil && cur.Val < tmp.Val {
			cur = cur.Next
			prev = prev.Next
		}

		prev.Next = tmp
		tmp.Next = cur
	}

	return vhead.Next
}
