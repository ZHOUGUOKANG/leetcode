package Num_327_CountRangeSum

import (
	"fmt"
	"testing"
)

func TestNum_327_CountRangeSum(t *testing.T) {

	e := []struct {
		nums         []int
		lower, upper int
		expected     int
	}{
		{[]int{-2, 5, -1}, -2, 2, 3},
	}
	for _, v := range e {
		o := countRangeSum(v.nums, v.lower, v.upper)
		if o != v.expected {
			t.Error(fmt.Sprintf("%v %v", v, o))
		}
	}
}
