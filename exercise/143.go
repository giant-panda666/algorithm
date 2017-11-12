package main

//Given a singly linked list L: L0→L1→…→Ln-1→Ln,
//reorder it to: L0→Ln→L1→Ln-1→L2→Ln-2→…
//
//You must do this in-place without altering the nodes' values.
//
//For example,
//Given {1,2,3,4}, reorder it to {1,4,2,3}.
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
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

func reorderList(head *ListNode) {
	if head == nil {
		return
	}

}
