package Num_101_IsSymmetric

import (
	"fmt"
	"testing"
)

func TestNum_101_IsSymmetric(t *testing.T) {
	e := []struct {
		nums     []int
		expected bool
	}{
		{[]int{1, 2, 2, 3, 4, 4, 3}, true},
		{[]int{1, 2, 2, -1, 3, -1, 3}, false},
		{[]int{1, 2, 2, 2, -1, 2}, false},
	}
	for _, v := range e {
		nodes := make([]*TreeNode, len(v.nums))
		for i, k := range v.nums {
			if k == -1 {
				nodes[i] = nil
			} else {
				nodes[i] = &TreeNode{Val: k}
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
		o := isSymmetric(nodes[0])
		if o != v.expected {
			t.Error(fmt.Sprintf("%v %v", v, o))
		}
	}
}
