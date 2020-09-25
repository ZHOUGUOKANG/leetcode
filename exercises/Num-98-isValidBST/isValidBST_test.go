package Num_98_IsValidBST

import (
	"fmt"
	"testing"
)

func TestNum_98_IsValidBST(t *testing.T) {

	e := []struct {
		nodes    []int
		expected bool
	}{
		{},
	}
	for _, v := range e {
		nodes := make([]*TreeNode, len(v.nodes))
		for i, val := range v.nodes {
			if val == -1 {
				nodes[i] = nil
			} else {
				nodes[i] = &TreeNode{Val: val}
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
		o := isValidBST(nodes[0])
		if o != v.expected {
			t.Error(fmt.Sprintf("%v %v", v, o))
		}
	}
}
