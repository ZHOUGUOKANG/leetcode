package Num_19_removeNthFromEnd

import (
	"reflect"
	"testing"
)

type testRemove struct {
	nums     []int
	n        int
	expected []int
}

func TestRemoveNthFromEnd(t *testing.T) {
	tests := []*testRemove{
		{[]int{1, 2, 3, 4, 5}, 2, []int{1, 2, 3, 5}},
		{[]int{1}, 1, []int{}},
		{[]int{1, 2}, 1, []int{1}},
		{[]int{1, 2}, 2, []int{2}},
	}
	for _, v := range tests {
		head := &ListNode{Val: v.nums[0]}
		tmp := head
		for i := 1; i < len(v.nums); i++ {
			tmp.Next = &ListNode{Val: v.nums[i]}
			tmp = tmp.Next
		}
		h := removeNthFromEnd(head, v.n)
		var values = []int{}
		for h != nil {
			values = append(values, h.Val)
			h = h.Next
		}
		if !reflect.DeepEqual(values, v.expected) {
			t.Error(v, values)
		}
	}
}
