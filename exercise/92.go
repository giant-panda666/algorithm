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

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var pm, pmPrev *ListNode = head, nil
	var pn, pnNext *ListNode = head, head.Next

	// find m node and its previous node
	for i := 0; i < m-1 && pm != nil; i++ {
		pmPrev = pm
		pm = pm.Next
	}
	// find n node and its next node
	for i := 0; i < n-1 && pnNext != nil; i++ {
		pn = pnNext
		pnNext = pnNext.Next
	}
	pn.Next = nil

	// reverse pm and pn
	tmpHead := reverse(pm, pn)

	if pmPrev != nil {
		pmPrev.Next = tmpHead
	} else {
		head = tmpHead
	}
	pm.Next = pnNext

	return head
}

func reverse(pm, pn *ListNode) *ListNode {
	var head, cur *ListNode = nil, pm
	for cur != nil {
		next := cur.Next
		cur.Next = head
		head = cur
		cur = next
	}
	return head
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
