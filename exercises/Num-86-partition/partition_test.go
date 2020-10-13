package Num_86_Partition

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNum_86_Partition(t *testing.T) {
	e := []struct {
		nums     []int
		x        int
		expected []int
	}{
		{[]int{1, 4, 3, 2, 5, 2}, 3, []int{1, 2, 2, 4, 3, 5}},
	}
	for _, v := range e {
		nodes := make([]*ListNode, len(v.nums))
		for k, val := range v.nums {
			nodes[k] = &ListNode{Val: val}
			if k > 0 {
				nodes[k-1].Next = nodes[k]
			}
		}
		o := partition(nodes[0], v.x)
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
