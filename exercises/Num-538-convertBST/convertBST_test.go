package Num_538_ConvertBST

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNum_538_ConvertBST(t *testing.T) {

	e := []struct {
		nodes    []int
		expected []int
	}{
		{},
	}
	for _, v := range e {
		nodes := make([]*TreeNode, len(v.nodes))
		for k, val := range v.nodes {
			if val == -1 {
				nodes[k] = nil
			} else {
				nodes[k] = &TreeNode{Val: val}
			}
			parent := nodes[(k-1)/2]
			if parent != nil {
				if k%2 == 0 {
					parent.Right = nodes[k]
				} else {
					parent.Left = nodes[k]
				}
			}

		}
		o := convertBST(nodes[0])
		ret := make([]int, 0)
		queue := []*TreeNode{o}
		for len(queue) > 0 {
			h := queue[0]
			queue = queue[1:]
			if h == nil {
				ret = append(ret, -1)
				continue
			}
			ret = append(ret, h.Val)
			queue = append(queue, h.Left, h.Right)
		}
		if !reflect.DeepEqual(v.expected, ret) {
			t.Error(fmt.Sprintf("%v %v", v, ret))
		}
	}
}
