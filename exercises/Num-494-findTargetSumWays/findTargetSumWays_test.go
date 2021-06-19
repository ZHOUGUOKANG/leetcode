package Num_494_FindTargetSumWays

import (
	"fmt"
	"testing"
)

func TestNum_494_FindTargetSumWays(t *testing.T) {
	e := []struct {
		nums     []int
		target   int
		expected int
	}{
		{nums: []int{1, 1, 1, 1, 1}, target: 3, expected: 5},
		{nums: []int{1}, target: 1, expected: 1},
		{nums: []int{1, 0}, target: 1, expected: 2},
	}
	for _, v := range e {
		o := findTargetSumWays(v.nums, v.target)
		if o != v.expected {
			t.Error(fmt.Sprintf("%v %v", v, o))
		}
	}
}
