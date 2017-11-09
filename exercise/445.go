package main

import "fmt"

//You are given two non-empty linked lists representing two non-negative integers. The most significant digit comes first and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.
//
//You may assume the two numbers do not contain any leading zero, except the number 0 itself.
//
//Follow up:
//What if you cannot modify the input lists? In other words, reversing the lists is not allowed.
//
//Example:
//
//Input: (7 -> 2 -> 4 -> 3) + (5 -> 6 -> 4)
//Output: 7 -> 8 -> 0 -> 7

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var len1, len2 int
	var h1, h2 = l1, l2
	for h1 != nil {
		h1 = h1.Next
		len1++
	}
	for h2 != nil {
		h2 = h2.Next
		len2++
	}

	h1, h2 = l1, l2
	var head *ListNode
	if len1 > len2 {
		head = l1
		for i := 0; i < len1-len2; i++ {
			h1 = h1.Next
		}
		for h1 != nil {
			h1.Val += h2.Val
			h1 = h1.Next
			h2 = h2.Next
		}
	} else {
		head = l2
		for i := 0; i < len2-len1; i++ {
			h2 = h2.Next
		}
		for h2 != nil {
			h2.Val += h1.Val
			h1 = h1.Next
			h2 = h2.Next
		}
	}

	var prev, cur *ListNode = nil, head
	for cur != nil {
		if cur.Val >= 10 {
			if prev == nil {
				prev = &ListNode{
					Val:  cur.Val / 10,
					Next: cur,
				}
				cur.Val %= 10
				head = prev
				cur = prev
				prev = nil
			} else {
				prev.Val += cur.Val / 10
				cur.Val %= 10
				prev = nil
				cur = head
			}
		} else {
			prev = cur
			cur = cur.Next
		}
	}

	return head
}

func makeList(a []int) *ListNode {
	if len(a) == 0 {
		return nil
	}
	head := &ListNode{
		Val:  a[len(a)-1],
		Next: nil,
	}
	for i := len(a) - 2; i >= 0; i-- {
		tmp := &ListNode{
			Val:  a[i],
			Next: head,
		}
		head = tmp
	}
	return head
}

func PrintList(h *ListNode) {
	for h != nil {
		fmt.Print(h.Val)
		h = h.Next
	}
	fmt.Println()
}

func main() {
	var a1 = []int{2, 4, 3}
	var a2 = []int{5, 6, 4}
	h1, h2 := makeList(a1), makeList(a2)
	PrintList(h1)
	PrintList(h2)

	h := addTwoNumbers(h1, h2)
	PrintList(h)
}
