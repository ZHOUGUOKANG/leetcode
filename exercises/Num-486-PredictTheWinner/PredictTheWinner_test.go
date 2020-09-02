package Num_486_PredictTheWinner

import (
	"fmt"
	"testing"
)

func TestPredictTheWinner(t *testing.T) {
	example := []struct {
		nums     []int
		expected bool
	}{
		{[]int{1, 5, 2}, false},
		{[]int{1, 5, 233, 7}, true},
		{[]int{1, 2, 99}, true},
		{[]int{1000, 999, 999, 1000, 555, 400}, true},
	}

	for _, v := range example {
		o := PredictTheWinner(v.nums)
		if o != v.expected {
			t.Error(fmt.Sprintf("%v %v", v, o))
		}
	}
}
