package Num_114_flatten

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

func flatten(root *TreeNode) {
	for root != nil {
		if root.Left != nil {
			right := root.Right
			root.Left, root.Right = nil, root.Left
			if right != nil {
				cur := root.Right
				for cur.Right != nil {
					cur = cur.Right
				}
				cur.Right = right
			}
		}
		root = root.Right
	}
}
