package Num_968_MinCameraCover

import (
	"fmt"
	"testing"
)

func TestNum_968_MinCameraCover(t *testing.T) {
	e := []struct {
		nodes    []int
		expected int
	}{
		{[]int{1, 2, -1, 3, 4}, 1},
		{[]int{1, 2, -1, 3, -1, -1, -1, 4, -1, -1, -1, -1, -1, -1, -1, -1, 5}, 2},
		{[]int{1}, 1},
	}
	for _, v := range e {
		nodes := make([]*TreeNode, len(v.nodes))
		for k, val := range v.nodes {
			if val != -1 {
				nodes[k] = &TreeNode{val, nil, nil}
			} else {
				nodes[k] = nil
			}
			parent := (k - 1) / 2
			if k > 0 && nodes[parent] != nil {
				if k%2 == 0 {
					nodes[parent].Right = nodes[k]
				} else {
					nodes[parent].Left = nodes[k]
				}
			}
		}
		o := minCameraCover(nodes[0])
		if o != v.expected {
			t.Error(fmt.Sprintf("%v %v", v, o))
		}
	}
}
