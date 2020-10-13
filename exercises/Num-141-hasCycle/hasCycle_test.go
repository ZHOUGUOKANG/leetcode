package Num_141_HasCycle

import (
	"fmt"
	"testing"
)

func TestNum_141_HasCycle(t *testing.T) {
	e := []struct {
		values   []int
		pos      int
		expected bool
	}{
		{[]int{3, 2, 0, -4}, 1, true},
		{[]int{1, 2}, 0, true},
		{[]int{1}, -1, false},
	}
	for _, v := range e {
		nodes := make([]*ListNode, len(v.values))
		for k, v := range v.values {
			nodes[k] = &ListNode{Val: v}
			if k > 0 {
				nodes[k-1].Next = nodes[k]
			}
		}
		if v.pos != -1 {
			nodes[len(nodes)-1].Next = nodes[v.pos]
		}
		o := hasCycle(nodes[0])
		if o != v.expected {
			t.Error(fmt.Sprintf("%v %v", v, o))
		}
	}
}
