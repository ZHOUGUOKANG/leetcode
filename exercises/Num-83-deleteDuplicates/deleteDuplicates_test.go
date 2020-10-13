package Num_83_DeleteDuplicates

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNum_83_DeleteDuplicates(t *testing.T) {
	e := []struct {
		nums     []int
		expected []int
	}{
		{[]int{1, 1, 2, 2, 3}, []int{1, 2, 3}},
		{[]int{1, 1, 2, 3, 3}, []int{1, 2, 3}},
		{[]int{1, 2, 2}, []int{1, 2}},
	}
	for _, v := range e {
		nodes := make([]*ListNode, len(v.nums))
		for k, v := range v.nums {
			nodes[k] = &ListNode{Val: v}
			if k > 0 {
				nodes[k-1].Next = nodes[k]
			}
		}
		o := deleteDuplicates(nodes[0])
		ex := make([]int, 0)
		for o != nil {
			ex = append(ex, o.Val)
			o = o.Next
		}
		if !reflect.DeepEqual(v.expected, ex) {
			t.Error(fmt.Sprintf("%v %v", v, ex))
		}
	}
}
