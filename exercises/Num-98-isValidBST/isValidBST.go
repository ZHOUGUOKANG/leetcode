package Num_98_IsValidBST

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//func isValidBST(root *TreeNode) bool {
//	if root == nil {
//		return true
//	}
//	var helper func(n *TreeNode, lower, upper int) bool
//	helper = func(n *TreeNode, lower, upper int) bool {
//		if n == nil {
//			return true
//		}
//		if n.Val > lower && n.Val < upper {
//			return helper(n.Left, lower, n.Val) && helper(n.Right, n.Val, upper)
//		}
//		return false
//	}
//	return helper(root, math.MinInt64, math.MaxInt64)
//}
func isValidBST(root *TreeNode) bool {
	var stack []*TreeNode
	min := math.MinInt64
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Val <= min {
			return false
		}
		min = root.Val
		root = root.Right
	}
	return true
}
