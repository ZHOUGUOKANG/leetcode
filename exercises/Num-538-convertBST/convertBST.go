package Num_538_ConvertBST

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func convertBST(root *TreeNode) *TreeNode {
	var visit func(n *TreeNode)
	sum := 0
	visit = func(n *TreeNode) {
		if n == nil {
			return
		}
		visit(n.Right)
		sum += n.Val
		n.Val = sum
		visit(n.Left)
	}
	visit(root)
	return root
}
