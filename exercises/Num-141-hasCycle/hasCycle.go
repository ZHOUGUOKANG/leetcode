package Num_141_HasCycle

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	l, r := head, head.Next
	for l != r {
		if r == nil || r.Next == nil {
			return false
		}
		l, r = l.Next, r.Next.Next
	}
	return true
}
