package num23

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
	//merge 2 lists
	lists[0] = merge2Lists(lists[0], lists[len(lists)-1])
	return mergeKLists(lists[:len(lists)-1])
}
func merge2Lists(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = merge2Lists(l1.Next, l2)
		return l1
	} else {
		l2.Next = merge2Lists(l1, l2.Next)
		return l2
	}
}
