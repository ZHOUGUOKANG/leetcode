package Num_109_sortedListToBST

type ListNode struct {
	Val  int
	Next *ListNode
}
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	var subList []int
	c := head
	for c != nil {
		subList = append(subList, c.Val)
		c = c.Next
	}
	return helper(subList)
}
func helper(subList []int) *TreeNode {
	l := len(subList)
	if l == 0 {
		return nil
	}
	midNum := l / 2
	root := &TreeNode{Val: subList[midNum]}
	root.Left = helper(subList[:midNum])
	root.Right = helper(subList[midNum+1:])
	return root
}
