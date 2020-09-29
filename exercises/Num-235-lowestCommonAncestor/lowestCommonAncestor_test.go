package Num_235_LowestCommonAncestor

import (
	"fmt"
	"testing"
)

func TestNum_235_LowestCommonAncestor(t *testing.T) {
	e := []struct {
		nums     []int
		p, q     int
		expected int
	}{
		{[]int{6, 2, 8, 0, 4, 7, 9, -1, -1, 3, 5}, 2, 8, 6},
		{[]int{6, 2, 8, 0, 4, 7, 9, -1, -1, 3, 5}, 2, 4, 2},
	}
	for _, v := range e {
		nodes := make([]*TreeNode, len(v.nums))
		var p, q, excepted *TreeNode
		for i, val := range v.nums {
			if val == -1 {
				nodes[i] = nil
			} else {
				nodes[i] = &TreeNode{Val: val}
				if val == v.p {
					p = nodes[i]
				}
				if val == v.q {
					q = nodes[i]
				}
				if val == v.expected {
					excepted = nodes[i]
				}
			}
			parent := nodes[(i-1)/2]
			if parent != nil {
				if i%2 == 0 {
					parent.Right = nodes[i]
				} else {
					parent.Left = nodes[i]
				}
			}

		}
		o := lowestCommonAncestor(nodes[0], p, q)
		if o != excepted {
			t.Error(fmt.Sprintf("%v %v", v, o))
		}
	}
}
