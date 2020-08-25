package Num_104_maxDepth

/**
* Definition for a binary tree node.
//* type TreeNode struct {
//*     Val int
//*     Left *TreeNode
//*     Right *TreeNode
//* }
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//func maxDepth(root *TreeNode) int {
//	if root == nil {
//		return 0
//	}
//	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
//}
//
//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	ans := 0
	for len(queue) > 0 {
		sz := len(queue)
		for sz > 0 {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			sz--
		}
		ans++
	}
	return ans
}
