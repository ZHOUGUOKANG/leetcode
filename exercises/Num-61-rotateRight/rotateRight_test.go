package Num_61_RotateRight

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNum_61_RotateRight(t *testing.T) {
	e := []struct {
		nums     []int
		k        int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, []int{4, 5, 1, 2, 3}},
		{[]int{0, 1, 2}, 4, []int{2, 0, 1}},
		{[]int{1, 2}, 2, []int{1, 2}},
		{[]int{1}, 99, []int{1}},
	}
	for _, v := range e {
		nodes := make([]*ListNode, len(v.nums))
		for k, val := range v.nums {
			nodes[k] = &ListNode{Val: val}
			if k > 0 {
				nodes[k-1].Next = nodes[k]
			}
		}
		o := rotateRight(nodes[0], v.k)
		nums := make([]int, 0)
		for o != nil {
			nums = append(nums, o.Val)
			o = o.Next
		}
		if !reflect.DeepEqual(nums, v.expected) {
			t.Error(fmt.Sprintf("%v %v", v, nums))
		}
	}
}
