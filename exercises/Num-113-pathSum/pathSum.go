package Num_113_PathSum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) [][]int {
	ans := make([][]int, 0)
	tmp := make([]int, 0)
	helper := func(node *TreeNode, tmpSum int) {}
	helper = func(node *TreeNode, tmpSum int) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil {
			if tmpSum+node.Val == sum {
				as := make([]int, len(tmp)+1)
				copy(as, tmp)
				as[len(as)-1] = node.Val
				ans = append(ans, as)
			}
		} else {
			tmpSum += node.Val
			tmp = append(tmp, node.Val)
			helper(node.Left, tmpSum)
			helper(node.Right, tmpSum)
			tmp = tmp[:len(tmp)-1]
		}
	}
	helper(root, 0)
	return ans
}
