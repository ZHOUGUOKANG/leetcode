package Num_142_DetectCycle

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for {
		if fast == nil || fast.Next == nil {
			return nil
		}
		slow, fast = slow.Next, fast.Next.Next
		if fast == slow {
			p := head
			for p != slow {
				p, slow = p.Next, slow.Next
			}
			return p
		}
	}
}
