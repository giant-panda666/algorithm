package main

import "fmt"

//Given a singly linked list L: L0→L1→…→Ln-1→Ln,
//reorder it to: L0→Ln→L1→Ln-1→L2→Ln-2→…
//
//You must do this in-place without altering the nodes' values.
//
//For example,
//Given {1,2,3,4}, reorder it to {1,4,2,3}.
// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

type stack struct {
	data  []*ListNode
	count int
}

func (s *stack) push(l *ListNode) {
	s.data = append(s.data, l)
	s.count++
}

func (s *stack) pop() *ListNode {
	if s.count == 0 {
		return nil
	}
	s.count--
	res := s.data[s.count]
	s.data = s.data[:s.count]
	return res
}

func (s *stack) isEmpty() bool {
	return s.count == 0
}

func reorderList(head *ListNode) {
	if head == nil {
		return
	}

	var p, q = head, head
	for q.Next != nil && q.Next.Next != nil {
		p = p.Next
		q = q.Next.Next
	}

	var s stack
	for q := p.Next; q != nil; {
		tmp := q.Next
		q.Next = nil
		s.push(q)
		q = tmp
	}
	p.Next = nil

	var cur = head
	for !s.isEmpty() {
		tmp := s.pop()
		tmp.Next = cur.Next
		cur.Next = tmp

		cur = cur.Next.Next
	}

	return
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
	var count = 0
	for h != nil {
		fmt.Print(h.Val)
		h = h.Next
		count++
	}
	fmt.Println()
}

func main() {
	var a1 = []int{1, 2, 3, 4}
	h := makeList(a1)
	reorderList(h)
	PrintList(h)
}
