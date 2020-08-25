package Num_25_reverseKGroup

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	i, c := 0, head
	for c != nil && i < k {
		i++
		c = c.Next
	}
	if i < k {
		return head
	}
	cc, n := head, head.Next
	for j := 0; j < k-1; j++ {
		n.Next, cc, n = cc, n, n.Next
	}
	head.Next = reverseKGroup(c, k)
	return cc
}
