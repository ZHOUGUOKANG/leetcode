package Num_105_BuildTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//func buildTree(preorder []int, inorder []int) *TreeNode {
//	cache := make(map[int]int)
//	for k,v := range inorder{
//		cache[v] = k
//	}
//	var helper func(pl,pr,il,ir int)*TreeNode
//	helper = func(pl, pr, il, ir int)*TreeNode {
//		if pl >= pr{
//			return nil
//		}
//		root := &TreeNode{Val:preorder[pl]}
//		index := cache[root.Val]
//		pl++
//		root.Left = helper(pl,pl+index-il,il,index)
//		root.Right = helper(pl+index-il,pr,index+1,ir)
//		return root
//	}
//	return helper(0,len(preorder),0,len(inorder))
//}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			root := &TreeNode{Val: preorder[0]}
			root.Left = buildTree(preorder[1:i+1], inorder[:i])
			root.Right = buildTree(preorder[i+1:], inorder[i+1:])
			return root
		}
	}
	return nil
}
