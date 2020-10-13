package Num_86_Partition

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	small, big := &ListNode{}, &ListNode{}
	s, b := small, big
	for head != nil {
		if head.Val >= x {
			b.Next, b = head, head
		} else {
			s.Next, s = head, head
		}
		head = head.Next
	}
	s.Next, b.Next = big.Next, nil
	return small.Next
}
