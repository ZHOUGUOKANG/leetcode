package Num_993_IsCousins

import (
	"fmt"
	"math"
	"testing"
)

func TestNum_993_IsCousins(t *testing.T) {

	e := []struct {
		nodes    []int
		x, y     int
		expected bool
	}{
		{[]int{1, 2, 3, math.MinInt64, 4, math.MinInt64, 5}, 4, 5, true},
	}
	for _, v := range e {
		nodes := make([]*TreeNode, len(v.nodes))
		for i, item := range v.nodes {
			if item == math.MinInt64 {
				continue
			}
			nodes[i] = &TreeNode{Val: item}
			parent := nodes[(i-1)/2]
			if parent != nil {
				if i%2 == 0 {
					parent.Right = nodes[i]
				} else {
					parent.Left = nodes[i]
				}
			}
		}
		o := isCousins(nodes[0], v.x, v.y)
		if o != v.expected {
			t.Error(fmt.Sprintf("%v %v", v, o))
		}
	}
}
