package Num_501_FindMode

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNum_501_FindMode(t *testing.T) {

	e := []struct {
		nums     []int
		expected []int
	}{
		{[]int{1, -1, 2, -1, -1, 2}, []int{2}},
		{[]int{1, -1, 2}, []int{1, 2}},
	}
	for _, v := range e {
		nodes := make([]*TreeNode, len(v.nums))
		for i, val := range v.nums {
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
		o := findMode(nodes[0])
		if !reflect.DeepEqual(v.expected, o) {
			t.Error(fmt.Sprintf("%v %v", v, o))
		}
	}
}
