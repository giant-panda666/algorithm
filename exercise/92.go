package main

import "fmt"

//Reverse a linked list from position m to n. Do it in-place and in one-pass.
//
//For example:
//Given 1->2->3->4->5->NULL, m = 2 and n = 4,
//
//return 1->4->3->2->5->NULL.
//
//Note:
//Given m, n satisfy the following condition:
//1 ≤ m ≤ n ≤ length of list.

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m, n int) *ListNode {
	if m == n {
		return head
	}

	var prevM = new(ListNode)
	newHead := prevM
	prevM.Next = head
	for i := 1; i < m; i++ {
		prevM = prevM.Next
	}
	tail := prevM.Next
	var prev, cur *ListNode = nil, prevM.Next
	for i := 0; i <= n-m; i++ {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	prevM.Next = prev
	tail.Next = cur
	return newHead.Next
}

func main() {
	var head = new(ListNode)
	head.Val = 3
	var next = new(ListNode)
	next.Val = 5
	head.Next = next

	res := reverseBetween(head, 1, 2)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}
