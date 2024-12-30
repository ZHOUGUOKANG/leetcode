package num19

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	root := &ListNode{Next: head}
	var i int
	var helper func(p, c *ListNode)
	helper = func(p, c *ListNode) {
		if c == nil {
			return
		}
		helper(c, c.Next)
		i++
		//delete self
		if n == i {
			p.Next = c.Next
			return
		}
	}
	helper(root, root.Next)
	return root.Next
}
