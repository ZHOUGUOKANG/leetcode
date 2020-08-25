package Num_111_minDepth

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l, r := minDepth(root.Left), minDepth(root.Right)
	if l == 0 || r == 0 {
		return max(l, r) + 1
	}
	return min(l, r) + 1
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
