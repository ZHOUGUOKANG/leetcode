package Num_257_binaryTreePaths

import (
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	temp := make([]int, 0)
	ans := make([]string, 0)
	bfs := func(node *TreeNode) {}
	bfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		temp = append(temp, node.Val)
		if node.Left == nil && node.Right == nil {
			str := strconv.Itoa(temp[0])
			for i := 1; i < len(temp); i++ {
				str += "->" + strconv.Itoa(temp[i])
			}
			ans = append(ans, str)
		} else {
			bfs(node.Left)
			bfs(node.Right)
		}
		temp = temp[:len(temp)-1]
	}
	bfs(root)
	return ans
}
