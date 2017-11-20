package main

//Merge k sorted linked lists and return it as one sorted list. Analyze and describe its complexity.

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type label struct {
	p     *ListNode
	isNil bool
}

func mergeKLists(lists []*ListNode) *ListNode {
	var vhead = new(ListNode)
	var lp []label
	var counter int
	for _, v := range lists {
		var tmp label
		tmp.p = v
		if v != nil {
			tmp.isNil = false
			counter++
		} else {
			tmp.isNil = true
		}
		lp = append(lp, tmp)
	}

	var min *ListNode
	for counter > 0 {
		for _, v := range lp {
			if v.isNil {
				continue
			}
			if min == nil {
				min = v
			} else if v.p.Val < min.Val {
				min = v

			}

		}

	}
}
