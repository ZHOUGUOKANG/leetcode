package Num_637_averageOfLevels

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type Level struct {
	nums  int
	count int
}

func averageOfLevels(root *TreeNode) []float64 {
	var vals []*Level
	helper := func(n *TreeNode, level int) {}
	helper = func(n *TreeNode, level int) {
		if n == nil {
			return
		}
		if level >= len(vals) {
			vals = append(vals, &Level{1, n.Val})
		} else {
			vals[level].count += n.Val
			vals[level].nums++
		}
		level++
		helper(n.Left, level)
		helper(n.Right, level)
	}
	helper(root, 0)
	ans := make([]float64, len(vals))
	for k, v := range vals {
		ans[k] = float64(v.count) / float64(v.nums)
	}
	return ans
}
