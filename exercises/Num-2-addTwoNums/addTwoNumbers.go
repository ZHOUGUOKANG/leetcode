package Num_2_addTwoNums

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	returnList := &ListNode{}
	num := 0
	carry := 0
	jumpList := returnList
	first := true
	for {
		num = l1.Val + l2.Val + carry
		carry = num / 10
		if first {
			jumpList.Val = num % 10
			first = false
		} else {
			jumpList.Next = &ListNode{Val: num % 10}
			jumpList = jumpList.Next
		}
		l1 = l1.Next
		l2 = l2.Next
		if l1 == nil && l2 == nil {
			if carry != 0 {
				jumpList.Next = &ListNode{Val: carry}
			}
			break
		}
		if l1 == nil {
			l1 = &ListNode{Val: 0}
		}
		if l2 == nil {
			l2 = &ListNode{Val: 0}
		}
	}
	return returnList
}
