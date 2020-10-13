package Num_530_GetMinimumDifference

import (
	"fmt"
	"testing"
)

func TestNum_530_GetMinimumDifference(t *testing.T) {
	e := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, -1, 3, -1, -1, 2}, 1},
		{[]int{5, 4, 7}, 1},
		{[]int{1, -1, 5, -1, -1, 3}, 2},
		{[]int{236, 104, 701, -1, 227, -1, 911}, 9},
		{[]int{600, 424, 612, -1, 499, -1, 689}, 12},
		{[]int{0, -1, 2236, -1, -1, 1277, 2776, -1, -1, -1, -1, 519}, 519},
	}
	for _, v := range e {
		nodes := make([]*TreeNode, len(v.nums))
		for k, val := range v.nums {
			if val != -1 {
				nodes[k] = &TreeNode{Val: val}
			} else {
				nodes[k] = nil
			}
			parent := nodes[(k-1)/2]
			if parent != nil {
				if k%2 == 0 {
					parent.Right = nodes[k]
				} else {
					parent.Left = nodes[k]
				}
			} else {
				if nodes[k] != nil {
					t.Error("make nodes err")
				}
			}
		}
		o := getMinimumDifference(nodes[0])
		if o != v.expected {
			t.Error(fmt.Sprintf("%v %v", v, o))
		}
	}
}
