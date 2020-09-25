package Num_589_Preorder

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	var ans []int
	helper := func(n *Node) {}
	helper = func(n *Node) {
		if n == nil {
			return
		}
		ans = append(ans, n.Val)
		for _, v := range n.Children {
			helper(v)
		}
	}
	helper(root)
	return ans
}
