package main

import (
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	l13 := &ListNode{Val: 4, Next: nil}
	l12 := &ListNode{Val: 2, Next: l13}
	l11 := &ListNode{Val: 1, Next: l12}

	l23 := &ListNode{Val: 4, Next: nil}
	l22 := &ListNode{Val: 3, Next: l23}
	l21 := &ListNode{Val: 1, Next: l22}

	got := mergeTwoLists(l11, l21)
	wanted := []int{1, 1, 2, 3, 4, 4}

	for _, want := range wanted {
		if got == nil {
			t.Fatal("end of linked list, expected more nodes")
		}
		// fmt.Println("got: ", got)
		if got.Val != want {
			t.Errorf("got %d, want %d", got.Val, want)
		}
		got = got.Next
	}
}
