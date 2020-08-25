package Num_173_BSTIterator

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type BSTIterator struct {
	root     *TreeNode
	nodeList []*TreeNode
	offset   int
}

func (this *BSTIterator) appendNode(node *TreeNode) {
	if node == nil {
		return
	}
	if node.Left != nil {
		this.appendNode(node.Left)
	}
	this.nodeList = append(this.nodeList, node)
	if node.Right != nil {
		this.appendNode(node.Right)
	}
}

func Constructor(root *TreeNode) BSTIterator {
	nodeList := make([]*TreeNode, 0)
	bst := BSTIterator{root, nodeList, 0}
	bst.appendNode(root)
	for key, value := range bst.nodeList {
		println(key, value.Val)
	}
	return bst
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	defer func() {
		this.offset++
	}()
	return this.nodeList[this.offset].Val
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return this.offset < len(this.nodeList)
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
