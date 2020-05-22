package leetcodes

/**
2. 两数相加
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

循环遍历节点，活动指针位移，记得进位
示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807
*/
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
