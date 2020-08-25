package Num_34_searchRange

import (
	"reflect"
	"testing"
)

func TestSearchRange(t *testing.T) {
	tests := []*struct {
		nums     []int
		target   int
		expected []int
	}{
		{[]int{5, 7, 7, 8, 8, 10}, 8, []int{3, 4}},
		{[]int{5, 7, 7, 8, 8, 10}, 6, []int{-1, -1}},
	}
	for _, v := range tests {
		out := searchRange(v.nums, v.target)
		if !reflect.DeepEqual(out, v.expected) {
			t.Error(v, out)
		}
	}
}
