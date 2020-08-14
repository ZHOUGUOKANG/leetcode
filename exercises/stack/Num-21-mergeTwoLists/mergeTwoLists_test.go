package Num_21_mergeTwoLists

import (
	"reflect"
	"testing"
)

type testMergeTwoLists struct {
	list1    []int
	list2    []int
	expected []int
}

func TestMergeTwoLists(t *testing.T) {
	tests := []testMergeTwoLists{
		{[]int{1, 2, 4}, []int{1, 3, 4}, []int{1, 1, 2, 3, 4, 4}},
		{[]int{1, 1}, []int{1}, []int{1, 1, 1}},
		{[]int{}, []int{}, []int{}},
	}
	for _, v := range tests {
		list1 := &ListNode{}
		l1 := list1
		for _, val := range v.list1 {
			l1.Next = &ListNode{Val: val}
			l1 = l1.Next
		}
		list2 := &ListNode{}
		l2 := list2
		for _, val := range v.list2 {
			l2.Next = &ListNode{Val: val}
			l2 = l2.Next
		}
		out := mergeTwoLists(list1.Next, list2.Next)
		r := []int{}
		for out != nil {
			r = append(r, out.Val)
			out = out.Next
		}
		if !reflect.DeepEqual(r, v.expected) {
			t.Error(v, r)
		}
	}
}
