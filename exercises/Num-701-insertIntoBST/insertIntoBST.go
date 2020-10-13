package Num_701_InsertIntoBST

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	t := root
	for {
		if val > t.Val {
			if t.Right == nil {
				t.Right = &TreeNode{Val: val}
				break
			}
			t = t.Right
		} else {
			if t.Left == nil {
				t.Left = &TreeNode{Val: val}
				break
			}
			t = t.Left
		}
	}
	return root
}
