package Num_110_isBalanced

import (
	"math"
	"testing"
)

type test struct {
	values   []int
	expected bool
}

func TestIsBalanced(t *testing.T) {
	tests := []*test{
		{[]int{3, 9, 20, math.MinInt32, math.MinInt32, 15, 7}, true},
		{expected: true},
		{[]int{1, 2, 2, 3, 3, math.MinInt32, math.MinInt32, 4, 4}, false},
		//{},
	}
	for _, v := range tests {
		out := true
		if len(v.values) == 0 {
			out = isBalanced(nil)
		} else {
			root := &TreeNode{Val: v.values[0]}
			nodes := []*TreeNode{root}
			for i := 1; i < len(v.values); i++ {
				val := v.values[i]
				if val == math.MinInt32 {
					nodes = append(nodes, nil)
					continue
				}
				node := &TreeNode{Val: val}
				nodes = append(nodes, node)
				parent := nodes[(i-1)/2]
				if i%2 == 0 {
					//println(parent.Val,"right",node.Val)
					parent.Right = node
				} else {
					//println(parent.Val,"left",node.Val)
					parent.Left = node
				}
			}
			out = isBalanced(root)
		}
		if out != v.expected {
			t.Error(v, out)
		}
	}
}
