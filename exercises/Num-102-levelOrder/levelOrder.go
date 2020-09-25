package Num_102_LevelOrder

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	helper := func(n *TreeNode, level int) {}
	helper = func(n *TreeNode, level int) {
		if n == nil {
			return
		}
		if len(ans) <= level {
			ans = append(ans, []int{n.Val})
		} else {
			ans[level] = append(ans[level], n.Val)
		}
		level++
		helper(n.Left, level)
		helper(n.Right, level)
	}
	helper(root, 0)
	return ans
}
