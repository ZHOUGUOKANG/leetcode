package Num_100_isSameTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil {
		if q == nil {
			return true
		} else {
			return false
		}
	}
	if q == nil {
		return false
	}
	if q.Val != p.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
