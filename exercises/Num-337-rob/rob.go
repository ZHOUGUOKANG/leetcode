package Num_337_rob

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	return max(helper(root))
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func helper(node *TreeNode) (int, int) {
	if node == nil {
		return 0, 0
	}
	left0, left1 := helper(node.Left)
	right0, right1 := helper(node.Right)
	return max(left0, left1) + max(right0, right1), left0 + right0 + node.Val
}
