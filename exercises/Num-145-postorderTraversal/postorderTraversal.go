package Num_145_PostorderTraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	var ans []int
	helper := func(n *TreeNode) {}
	helper = func(n *TreeNode) {
		if n == nil {
			return
		}
		helper(n.Left)
		helper(n.Right)
		ans = append(ans, n.Val)
	}
	helper(root)
	return ans
}
