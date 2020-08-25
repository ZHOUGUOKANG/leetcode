package Num_23_mergeKLists

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

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	list1 := mergeKLists(lists[:len(lists)/2])
	list2 := mergeKLists(lists[len(lists)/2:])
	return mergeTwoLists(list1, list2)
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
