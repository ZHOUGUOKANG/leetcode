package num25

type ListNode struct {
	Val  int
	Next *ListNode
}

//func reverseKGroup(head *ListNode, k int) *ListNode {
//	head = &ListNode{Next: head}
//	cur, last := head.Next, head
//	stack := make([]*ListNode, 0, k)
//	for cur != nil {
//		last.Next = cur
//		for ; cur != nil && len(stack) < k; cur = cur.Next {
//			stack = append(stack, cur)
//		}
//		if len(stack) != k {
//			break
//		}
//		last.Next, last, stack[0].Next = stack[len(stack)-1], stack[0], cur
//		for len(stack) > 1 {
//			stack[len(stack)-1].Next = stack[len(stack)-2]
//			stack = stack[:len(stack)-1]
//		}
//		stack = stack[:0]
//	}
//	return head.Next
//}

func reverseKGroup(head *ListNode, k int) *ListNode {
	d := &ListNode{}
	p, c := d, head
	for c != nil {
		s := c
		for i := 0; i < k; i++ {
			if c == nil {
				p.Next = s
				return d.Next
			}
			c = c.Next
		}
		nextP := s
		for per, next := s, s.Next; next != c; {
			per.Next, s, next.Next, next = next.Next, next, s, next.Next
		}
		p.Next, p = s, nextP
	}
	return d.Next
}
