package Num_103_ZigzagLevelOrder

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	queue := []*TreeNode{root}
	flag := true
	var ans [][]int
	var pop *TreeNode
	for len(queue) > 0 {
		l := len(queue)
		res := make([]int, l)
		for i := 0; i < l; i++ {
			pop = queue[i]
			if flag {
				res[i] = pop.Val
			} else {
				res[l-i-1] = pop.Val
			}
			if pop.Left != nil {
				queue = append(queue, pop.Left)
			}
			if pop.Right != nil {
				queue = append(queue, pop.Right)
			}
		}
		queue = queue[l:]
		ans = append(ans, res)
		flag = !flag
	}
	return ans
}

//func zigzagLevelOrder(root *TreeNode) [][]int {
//	ans := make([][]int, 0)
//	helper := func(n *TreeNode, level int) {}
//	helper = func(n *TreeNode, level int) {
//		if n == nil {
//			return
//		}
//		if len(ans) <= level {
//			ans = append(ans, []int{n.Val})
//		} else {
//			ans[level] = append(ans[level], n.Val)
//		}
//		level++
//		helper(n.Left, level)
//		helper(n.Right, level)
//	}
//	helper(root, 0)
//	for i := 0; i < len(ans); i++ {
//		if i%2 != 0 {
//			levelLen := len(ans[i])
//			for j := 0; j < levelLen/2; j++ {
//				ans[i][j], ans[i][levelLen-1-j] = ans[i][levelLen-1-j], ans[i][j]
//			}
//		}
//	}
//	return ans
//}
