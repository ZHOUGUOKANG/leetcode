package Num_19_removeNthFromEnd

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	tmpNode := &ListNode{Next: head}
	left, right := tmpNode, tmpNode
	index := -1
	for right != nil {
		right = right.Next
		index++
		if index > n {
			left = left.Next
		}
		if right != nil {
			println(index, "left", left.Val, "right", right.Val)
		}
	}
	if left == tmpNode {
		return head.Next
	} else if left.Next != nil {
		println(left.Val)
		left.Next = left.Next.Next
	} else {
		left.Next = nil
	}
	return head
}
