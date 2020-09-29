package helper

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BuildTree(nums []int) interface{} {
	nodes := make([]*TreeNode, len(nums))
	for k, v := range nums {
		if v == -1 {
			nodes[k] = nil
		} else {
			nodes[k] = &TreeNode{Val: v}
		}
		parent := nodes[(k-1)/2]
		if parent != nil && k > 0 {
			if k%2 == 0 {
				parent.Right = nodes[k]
			} else {
				parent.Left = nodes[k]
			}
		}
	}
	return nodes[0]
}

func PreorderTraversal(root *TreeNode) []int {
	ans := make([]int, 0)
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		pop := stack[0]
		stack = stack[1:]
		if pop == nil {
			ans = append(ans, -1)
			continue
		}
		ans = append(ans, pop.Val)
		stack = append(stack, pop.Left, pop.Right)
	}
	return ans
}
