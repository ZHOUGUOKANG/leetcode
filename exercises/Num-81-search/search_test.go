package Num_81_Search

import (
	"fmt"
	"testing"
)

func TestNum_81_Search(t *testing.T) {
	e := []struct {
		nums     []int
		target   int
		expected bool
	}{
		{[]int{2, 5, 6, 0, 0, 1, 2}, 0, true},
		{[]int{2, 5, 6, 0, 0, 1, 2}, 3, false},
		{[]int{2, 5, 6, 0, 0, 1, 2, 2}, 6, true},
	}
	for _, v := range e {
		o := search(v.nums, v.target)
		if o != v.expected {
			t.Error(fmt.Sprintf("%v %v", v, o))
		}
	}
}
