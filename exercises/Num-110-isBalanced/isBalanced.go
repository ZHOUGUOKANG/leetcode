package Num_110_isBalanced

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

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return helper(root) != -1
}
func helper(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l, r := helper(root.Left), helper(root.Right)
	if l == -1 || r == -1 {
		return -1
	}
	h := l - r
	if -1 <= h && h <= 1 {
		return max(l, r) + 1
	} else {
		return -1
	}
}

//func height(root *TreeNode)int {
//	if root == nil{
//		return 0
//	}
//	return max(height(root.Left),height(root.Right)) + 1
//}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
