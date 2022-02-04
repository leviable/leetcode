package main

import (
	"sync"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

const MAX = 1_000_000_000

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	var start, prev, head *ListNode

	var v1, v2 int
	var n1, n2 *ListNode

	prev = &ListNode{}
	var once sync.Once
	for {
		if list1 == nil && list2 == nil {
			break
		}

		head = &ListNode{}
		prev.Next = head

		once.Do(func() { start = head })

		if list1 != nil {
			v1 = list1.Val
			n1 = list1.Next
		} else {
			v1 = MAX
			n1 = nil
		}

		if list2 != nil {
			v2 = list2.Val
			n2 = list2.Next
		} else {
			v2 = MAX
			n2 = nil
		}

		if v1 < v2 {
			head.Val = v1
			list1 = n1
		} else {
			head.Val = v2
			list2 = n2
		}

		prev = head
	}

	return start
}
