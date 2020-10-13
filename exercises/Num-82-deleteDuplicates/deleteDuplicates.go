package Num_82_DeleteDuplicates

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	t := head.Next
	for t != nil && t.Val == head.Val {
		t = t.Next
	}
	if head.Next != t {
		return deleteDuplicates(t)
	}
	head.Next = deleteDuplicates(t)
	return head
}
