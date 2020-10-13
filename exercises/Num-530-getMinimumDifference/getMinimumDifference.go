package Num_530_GetMinimumDifference

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	ans, pre := math.MaxInt64, -1
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if pre != -1 && node.Val-pre < ans {
			ans = node.Val - pre
		}
		pre = node.Val
		dfs(node.Right)
	}
	dfs(root)
	return ans
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//func getMinimumDifference(root *TreeNode) int {
//	if root == nil {
//		return 0
//	}
//	return helper(root)
//}

//func helper(root *TreeNode) int {
//	if root == nil {
//		return math.MaxInt64
//	}
//	minVal := math.MaxInt64
//	var dfs func(n *TreeNode,left bool)int
//	dfs = func(n *TreeNode,left bool) int {
//		if n == nil{
//			return math.MaxInt64
//		}
//		if left{
//			for n.Right != nil{
//				n = n.Right
//			}
//			return n.Val
//		}else {
//			for n.Left != nil{
//				n = n.Left
//			}
//			return n.Val
//		}
//	}
//
//	minVal = min(abs(root.Val-dfs(root.Left,true)),abs(root.Val-dfs(root.Right,false)))
//	return min(minVal,min(helper(root.Left), helper(root.Right)))
//}
