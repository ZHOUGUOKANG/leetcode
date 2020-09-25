package Num_501_FindMode

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMode(root *TreeNode) []int {
	maxCount, base, count := 0, math.MaxInt64, 0
	ans := make([]int, 0)
	helper := func(n *TreeNode) {}
	helper = func(n *TreeNode) {
		if n == nil {
			return
		}
		helper(n.Left)
		if n.Val == base {
			count++
		} else {
			if count > maxCount {
				maxCount = count
				if len(ans) > 0 {
					ans[0] = base
					ans = ans[:1]
				} else {
					ans = append(ans, base)
				}
			} else if count == maxCount {
				ans = append(ans, base)
			}
			base, count = n.Val, 1
		}
		helper(n.Right)
	}
	helper(root)
	if count > 0 {
		if count > maxCount {
			ans = []int{base}
		} else if count == maxCount {
			ans = append(ans, base)
		}
	}
	return ans
}
