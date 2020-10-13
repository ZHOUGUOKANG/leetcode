package Num_83_DeleteDuplicates

type ListNode struct {
	Val  int
	Next *ListNode
}

//func deleteDuplicates(head *ListNode) *ListNode {
//	if head == nil{
//		return head
//	}
//	var helper func(n *ListNode,pNum int)*ListNode
//	helper = func(n *ListNode, pNum int)*ListNode {
//		if n == nil{
//			return nil
//		}
//		if n.Val == pNum{
//			return helper(n.Next,pNum)
//		}
//		n.Next = helper(n.Next,n.Val)
//		return n
//	}
//	head.Next = helper(head.Next,head.Val)
//	return head
//}

//func deleteDuplicates(head *ListNode) *ListNode {
//	if head == nil {
//		return head
//	}
//	pNum, cur, last := head.Val, head.Next, head
//	for cur != nil {
//		if cur.Val == pNum {
//			for cur != nil && cur.Val == pNum {
//				cur = cur.Next
//			}
//			last.Next, last = cur, cur
//			if cur != nil {
//				pNum = cur.Val
//			}
//		} else {
//			pNum, cur, last = cur.Val, cur.Next, last.Next
//		}
//	}
//	return head
//}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	s, p := head, head.Next
	for p != nil {
		if s.Val == p.Val {
			s.Next, p = p.Next, p.Next
		} else {
			s, p = s.Next, p.Next
		}
	}
	return head
}
