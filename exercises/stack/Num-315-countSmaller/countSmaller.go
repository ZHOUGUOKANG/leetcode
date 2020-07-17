package Num_315_countSmaller

//func countSmaller(nums []int) []int {
//	length := len(nums)
//	result := make([]int,len(nums))
//	for index:=length-2;index>=0;index--{
//		for i:=index+1;i<length;i++  {
//			if nums[i] < nums[index] {
//				result[index]++
//			}else if nums[i] == nums[index] {
//				result[index] += result[i]
//				break
//			}
//
//		}
//	}
//	return result
//}

type TreeNode struct {
	Val   int
	Count int
	Left  *TreeNode
	Right *TreeNode
}

func insert(node, root *TreeNode, count *int) {
	if node.Val <= root.Val {
		root.Count++
		if root.Left != nil {
			insert(node, root.Left, count)
		} else {
			root.Left = node
		}
	}
	if node.Val > root.Val {
		*count = 1 + root.Count + *count
		if root.Right != nil {
			insert(node, root.Right, count)
		} else {
			root.Right = node
		}
	}
}

func countSmaller(nums []int) []int {
	length := len(nums)
	result := make([]int, len(nums))
	if length <= 1 {
		return result
	}
	root := &TreeNode{Val: nums[length-1]}
	var count int
	for index := length - 2; index >= 0; index-- {
		count = 0
		node := &TreeNode{Val: nums[index]}
		insert(node, root, &count)
		result[index] = count
	}
	return result
}
