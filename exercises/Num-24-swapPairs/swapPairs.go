package Num_24_swapPairs

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

func swapPairs(head *ListNode) *ListNode {
	s := &ListNode{Next: head}
	return swap(s)
}
func swap(head *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}
	head.Next, head.Next.Next, head.Next.Next.Next = head.Next.Next, head.Next.Next.Next, head.Next
	swap(head.Next.Next)
	return head.Next
}

/**
func swapPairs(head *ListNode) *ListNode {
	s := &ListNode{Next:head}
	helper := func(node *ListNode) {}
	helper = func(node *ListNode) {
		if node == nil || node.Next == nil || node.Next.Next == nil{
			return
		}
		node.Next,node.Next.Next,node.Next.Next.Next = node.Next.Next,node.Next.Next.Next,node.Next
		helper(node.Next.Next)
	}
	helper(s)
	return s.Next
}
*/
