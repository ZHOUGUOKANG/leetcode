package Num_147_InsertionSortList

import "math"

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	tempHead, sortedHead := &ListNode{Val: math.MinInt64, Next: head}, head
	for l, r := tempHead, head; r != nil; l, r = l.Next, r.Next {
		if r.Val < l.Val {
			sortedHead = r
		}
	}
	for t := head; t != sortedHead; t = tempHead.Next {
		tempHead.Next = t.Next
		if t.Val < sortedHead.Val {
			sortedHead.Next, t.Next, sortedHead.Val, t.Val = t, sortedHead.Next, t.Val, sortedHead.Val
		} else {
			for l, r := sortedHead, sortedHead.Next; ; l, r = l.Next, r.Next {
				if r == nil || r.Val > t.Val {
					l.Next, t.Next = t, r
					break
				}
			}
		}
	}
	return tempHead.Next
}
