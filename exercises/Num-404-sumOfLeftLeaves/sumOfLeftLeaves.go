package Num_404_SumOfLeftLeaves

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		return root.Left.Val + sumOfLeftLeaves(root.Right)
	} else {
		return sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)
	}
	//sum := 0
	//helper := func(node *TreeNode,left bool) {}
	//helper = func(node *TreeNode,left bool) {
	//	if node == nil{
	//		return
	//	}
	//	if left && node.Left == nil && node.Right == nil{
	//		sum += node.Val
	//		return
	//	}
	//	helper(node.Right,false)
	//	helper(node.Left,true)
	//}
	//helper(root,false)
	//return sum
}
