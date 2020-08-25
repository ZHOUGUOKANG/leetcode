package Num_21_mergeTwoLists

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

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	h := head
	t1, t2 := l1, l2
	for t1 != nil || t2 != nil {
		if t1 == nil {
			h.Next = t2
			break
		} else if t2 == nil {
			h.Next = t1
			break
		} else if t1.Val > t2.Val {
			h.Next, t2 = t2, t2.Next
		} else {
			h.Next, t1 = t1, t1.Next
		}
		h = h.Next
	}
	return head.Next
}

//func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
//	if l1 == nil{
//		return l2
//	}else if l2 == nil{
//		return l1
//	}else if l1.Val > l2.Val {
//		l2.Next = mergeTwoLists(l1,l2.Next)
//		return l2
//	}else {
//		l1.Next = mergeTwoLists(l1.Next,l2)
//		return l1
//	}
//}
