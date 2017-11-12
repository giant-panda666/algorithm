package main

//Remove all elements from a linked list of integers that have value val.
//
//Example
//Given: 1 --> 2 --> 6 --> 3 --> 4 --> 5 --> 6, val = 6
//Return: 1 --> 2 --> 3 --> 4 --> 5
func removeElements(head *ListNode, val int) *ListNode {
	var vhead = new(ListNode)
	vhead.Next = head
	var cur, del = vhead, head
	for del != nil {
		if del.Val == val {
			cur.Next = del.Next
			del.Next = nil
			del = cur.Next
		} else {
			cur = del
			del = del.Next
		}
	}

	return vhead.Next
}
