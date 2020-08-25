package Num_108_sortedArrayToBST

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	return createTree(nums, 0, len(nums)-1)
}

func createTree(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := (right + left) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = createTree(nums, left, mid-1)
	root.Right = createTree(nums, mid+1, right)
	return root
}
