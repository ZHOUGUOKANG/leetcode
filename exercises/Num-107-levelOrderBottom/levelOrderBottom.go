package Num_107_levelOrderBottom

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	var ret [][]int
	helper := func(node *TreeNode, level int) {}
	helper = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		if len(ret) < level+1 {
			ret = append(ret, []int{})
		}
		ret[level] = append(ret[level], node.Val)
		helper(node.Left, level+1)
		helper(node.Right, level+1)
	}
	helper(root, 0)
	for i := 0; i < len(ret)/2; i++ {
		ret[i], ret[len(ret)-i-1] = ret[len(ret)-i-1], ret[i]
	}
	return ret
}
