package Num_24_swapPairs

import (
	"reflect"
	"testing"
)

type test struct {
	nums     []int
	expected []int
}

func TestSwapPairs(t *testing.T) {
	tests := []*test{
		{[]int{1, 2, 3, 4}, []int{2, 1, 4, 3}},
		{[]int{1, 2, 3}, []int{2, 1, 3}},
	}
	for _, v := range tests {
		head := &ListNode{Val: v.nums[0]}
		c := head
		for i := 1; i < len(v.nums); i++ {
			c.Next = &ListNode{Val: v.nums[i]}
			c = c.Next
		}
		h := swapPairs(head)
		var out []int
		for h != nil {
			out = append(out, h.Val)
			h = h.Next
		}
		if !reflect.DeepEqual(out, v.expected) {
			t.Error(v, out)
		}
	}
}
