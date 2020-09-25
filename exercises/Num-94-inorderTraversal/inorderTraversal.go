package Num_94_inorderTraversal

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

//func inorderTraversal(root *TreeNode) []int {
//	result := make([]int,0)
//	if root == nil{
//		return result
//	}
//	result = append(result,inorderTraversal(root.Left)...)
//	result = append(result,root.Val)
//	result = append(result,inorderTraversal(root.Right)...)
//	return result
//}

//func inorderTraversal(root *TreeNode) []int {
//	result := make([]int, 0)
//	stack := make([]*TreeNode, 0)
//	current := root
//	for current != nil || len(stack) > 0 {
//		for current != nil {
//			stack = append(stack, current)
//			current = current.Left
//		}
//		current = stack[len(stack)-1]
//		stack = stack[:len(stack)-1]
//		result = append(result, current.Val)
//		current = current.Right
//	}
//	return result
//}

func inorderTraversal(root *TreeNode) []int {
	ans := make([]int, 0)
	helper := func(node *TreeNode) {}
	helper = func(node *TreeNode) {
		if node == nil {
			return
		}
		helper(node.Left)
		ans = append(ans, node.Val)
		helper(node.Right)
	}
	helper(root)
	return ans
}
