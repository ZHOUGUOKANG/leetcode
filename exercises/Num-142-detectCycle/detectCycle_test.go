package Num_142_DetectCycle

import (
	"fmt"
	"testing"
)

func TestNum_142_DetectCycle(t *testing.T) {
	e := []struct {
		nums []int
		pos  int
	}{
		{[]int{3, 2, 0, -4}, 1},
		{[]int{1, 2}, 0},
		{[]int{1}, -1},
		{[]int{-21, 10, 17, 8, 4, 26, 5, 35, 33, -7, -16, 27, -12, 6, 29, -12, 5, 9, 20, 14, 14, 2, 13, -24, 21, 23, -21, 5}, 24},
	}
	for _, v := range e {
		nodes := make([]*ListNode, len(v.nums))
		for k, val := range v.nums {
			nodes[k] = &ListNode{Val: val}
			if k > 0 {
				nodes[k-1].Next = nodes[k]
			}
		}
		if v.pos >= 0 {
			nodes[len(nodes)-1].Next = nodes[v.pos]
		}
		o := detectCycle(nodes[0])
		if v.pos == -1 {
			if o != nil {
				t.Error(fmt.Sprintf("%v %v", v, o))
			}
		} else {
			if o != nodes[v.pos] {
				t.Error(fmt.Sprintf("%v %v", v, o))
			}
		}

	}
}
