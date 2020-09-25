package Num_144_PreorderTraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	ans := make([]int, 0)
	helper := func(n *TreeNode) {}
	helper = func(n *TreeNode) {
		if n == nil {
			return
		}
		ans = append(ans, n.Val)
		helper(n.Left)
		helper(n.Right)
	}
	helper(root)
	return ans
}
