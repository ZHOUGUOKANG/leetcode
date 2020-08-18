package Num_27_removeElement

import (
	"reflect"
	"testing"
)

type test struct {
	nums         []int
	val          int
	expectedOut  int
	expectedNums []int
}

func TestAll(t *testing.T) {
	tests := []*test{
		{[]int{3, 2, 2, 3}, 3, 2, []int{2, 2}},
		{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 5, []int{0, 1, 4, 0, 3}},
		{[]int{2}, 3, 1, []int{2}},
		{[]int{3, 3}, 5, 2, []int{3, 3}},
		{[]int{3, 3}, 3, 0, []int{}},
	}
	for _, v := range tests {
		copyNums := make([]int, len(v.nums))
		copy(copyNums, v.nums)
		out := removeElement(v.nums, v.val)
		if out != v.expectedOut && !reflect.DeepEqual(v.expectedNums, v.nums[:out]) {
			t.Error(copyNums, v.expectedOut, v.nums[:out], out)
		}

	}
}
