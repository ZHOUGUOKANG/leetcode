package num24

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	head, head.Next.Next, head.Next = head.Next, head, head.Next.Next
	head.Next.Next = swapPairs(head.Next.Next)
	return head
}
