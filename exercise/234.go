package main

//Given a singly linked list, determine if it is a palindrome.
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	var p, q = head, head.Next
	for q != nil && q.Next != nil {
		p = p.Next
		q = q.Next.Next
	}

	var prev, cur *ListNode = p, p.Next
	p.Next = nil
	for cur != nil {
		tmp := cur.Next
		cur.Next = prev.Next
		prev.Next = cur
		cur = tmp
	}

	for l1, l2 := head, p.Next; l2 != nil; {
		if l1.Val != l2.Val {
			return false
		}

		l1 = l1.Next
		l2 = l2.Next
	}

	return true
}
