package Num_106_BuildTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	pLen, iLen := len(postorder), len(inorder)
	if pLen == 0 || iLen == 0 {
		return nil
	}
	tail := postorder[pLen-1]
	for k := 0; k < iLen; k++ {
		if inorder[k] == tail {
			node := &TreeNode{Val: inorder[k]}
			node.Left = buildTree(inorder[:k], postorder[:k])
			node.Right = buildTree(inorder[k+1:], postorder[k:pLen-1])
			return node
		}
	}
	return nil
}
