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

type stack struct {
	data  []*ListNode
	count int
}

func (s *stack) push(d *ListNode) {
	s.data = append(s.data[:s.count], d)
	s.count++
}

func (s *stack) pop() *ListNode {
	s.count--
	return s.data[s.count]
}

func (s *stack) isEmpty() bool {
	return s.count == 0
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var s1, s2 stack
	for l1 != nil {
		s1.push(l1)
		l1 = l1.Next
	}
	for l2 != nil {
		s2.push(l2)
		l2 = l2.Next
	}

	var sum = 0
	var head = new(ListNode)
	for !s1.isEmpty() || !s2.isEmpty() {
		if !s1.isEmpty() {
			tmp := s1.pop()
			sum += tmp.Val
		}
		if !s2.isEmpty() {
			tmp := s2.pop()
			sum += tmp.Val
		}
		list := new(ListNode)
		list.Val = sum % 10
		list.Next = head.Next
		head.Next = list
		sum /= 10
	}
	if sum != 0 {
		head.Val = sum
		return head
	}
	return head.Next
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
	var a1 = []int{0}
	var a2 = []int{7, 3}
	h1, h2 := makeList(a1), makeList(a2)
	PrintList(h1)
	PrintList(h2)

	h := addTwoNumbers(h1, h2)
	PrintList(h)
}
