package Num_61_RotateRight

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}
	l, r := head, head
	i := 1
	for r.Next != nil {
		r = r.Next
		i++
		if i > k+1 {
			l = l.Next
		}
	}
	if i == 1 || i == k {
		return head
	}
	if i < k {
		return rotateRight(head, k%i)
	}
	r.Next = head
	root := l.Next
	l.Next = nil
	return root
}
