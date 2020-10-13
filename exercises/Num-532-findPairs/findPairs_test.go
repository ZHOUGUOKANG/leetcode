package Num_532_FindPairs

import (
	"fmt"
	"testing"
)

func TestNum_532_FindPairs(t *testing.T) {
	e := []struct {
		nums     []int
		k        int
		expected int
	}{
		{[]int{3, 1, 4, 1, 5}, 2, 2},
		{[]int{3, 1, 4, 1, 5}, 0, 1},
		{[]int{1, 2, 3, 4, 5}, 1, 4},
		{[]int{1, 2, 4, 4, 3, 3, 0, 9, 2, 3}, 3, 2},
		{[]int{-1, -2, -3}, 1, 2},
		{[]int{1, 1, 1, 1, 1}, 0, 1},
	}
	for _, v := range e {
		o := findPairs(v.nums, v.k)
		if o != v.expected {
			t.Error(fmt.Sprintf("%v %v", v, o))
		}
	}
}
