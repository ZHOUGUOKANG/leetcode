package Num_31_nextPermutation

import (
	"reflect"
	"testing"
)

type test struct {
	nums     []int
	expected []int
}

func TestNextPermutation(t *testing.T) {
	tests := []*test{
		{[]int{1, 2, 3}, []int{1, 3, 2}},
		{[]int{1, 3, 2}, []int{2, 1, 3}},
		{[]int{3, 2, 1}, []int{1, 2, 3}},
		{[]int{1, 1, 5}, []int{1, 5, 1}},
		{[]int{1, 9, 0, 9}, []int{1, 9, 9, 0}},
	}
	for _, v := range tests {
		numsCopy := make([]int, len(v.nums))
		copy(numsCopy, v.nums)
		nextPermutation(numsCopy)
		if !reflect.DeepEqual(numsCopy, v.expected) {
			t.Error(v, v.nums)
		}
	}
}
