package Num_82_DeleteDuplicates

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNum_82_DeleteDuplicates(t *testing.T) {
	e := []struct {
		nums     []int
		expected []int
	}{
		{[]int{1, 2, 3, 3, 4, 4, 5}, []int{1, 2, 5}},
		{[]int{1, 1, 1, 2, 3}, []int{2, 3}},
		{[]int{1}, []int{1}},
		{[]int{1, 1}, []int{}},
	}
	for _, v := range e {
		nodes := make([]*ListNode, len(v.nums))
		for k, val := range v.nums {
			nodes[k] = &ListNode{Val: val}
			if k > 0 {
				nodes[k-1].Next = nodes[k]
			}
		}
		o := deleteDuplicates(nodes[0])
		expected := make([]int, 0)
		for o != nil {
			expected = append(expected, o.Val)
			o = o.Next
		}
		if !reflect.DeepEqual(v.expected, expected) {
			t.Error(fmt.Sprintf("%v %v", v, expected))
		}
	}
}
