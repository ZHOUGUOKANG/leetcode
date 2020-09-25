package Num_offer_04_05_isValidBST

import "math"

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

func isValidBST(root *TreeNode) bool {
	return helper(root, math.MinInt32-1, math.MaxInt32+1)
}

func helper(root *TreeNode, min, max int) (result bool) {
	if root == nil {
		return true
	}
	//println(root.Val,min,max)
	if root.Val <= min || root.Val >= max {
		return false
	}
	if root.Left != nil {
		if root.Val <= root.Left.Val || root.Left.Val >= max {
			return false
		}
	}
	if root.Right != nil {
		if root.Val >= root.Right.Val || root.Right.Val <= min {
			return false
		}
	}
	result = helper(root.Left, min, minValue(root.Val, max)) && helper(root.Right, maxValue(root.Val, min), max)
	//println(root.Val,result)
	return
}

func maxValue(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func minValue(x, y int) int {
	if x > y {
		return y
	}
	return x
}
