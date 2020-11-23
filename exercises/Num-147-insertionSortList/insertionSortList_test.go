package Num_147_InsertionSortList

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNum_147_InsertionSortList(t *testing.T) {

	e := []struct {
		nums     []int
		expected []int
	}{
		{[]int{4, 2, 1, 3}, []int{1, 2, 3, 4}},
		{[]int{-1, 5, 3, 4, 0}, []int{-1, 0, 3, 4, 5}},
	}
	for _, v := range e {
		nodes := make([]*ListNode, len(v.nums))
		for k, v := range v.nums {
			nodes[k] = &ListNode{Val: v}
			if k > 0 {
				nodes[k-1].Next = nodes[k]
			}
		}
		o := insertionSortList(nodes[0])
		e := make([]int, 0)
		for o != nil {
			e = append(e, o.Val)
			o = o.Next
		}
		if !reflect.DeepEqual(v.expected, e) {
			t.Error(fmt.Sprintf("%v %v", v, e))
		}
	}
}
