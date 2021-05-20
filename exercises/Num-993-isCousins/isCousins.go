package Num_993_IsCousins

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCousins(root *TreeNode, x int, y int) bool {
	var parentX, parentY *TreeNode
	var depthX, depthY int
	var foundX, foundY bool
	var helper func(parent *TreeNode, current *TreeNode, depth int)
	helper = func(parent *TreeNode, current *TreeNode, depth int) {
		if current == nil {
			return
		}
		if current.Val == x {
			parentX = parent
			depthX = depth
			foundX = true
		} else if current.Val == y {
			parentY = parent
			depthY = depth
			foundY = true
		}
		if foundX && foundY {
			return
		}
		helper(current, current.Left, depth+1)
		helper(current, current.Right, depth+1)
	}
	helper(nil, root, 0)
	return foundX && foundY && (depthX == depthY) && (parentX != parentY)
}
