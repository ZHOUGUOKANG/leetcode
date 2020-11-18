package Num_134_CanCompleteCircuit

import (
	"fmt"
	"testing"
)

func TestNum_134_CanCompleteCircuit(t *testing.T) {
	e := []struct {
		gas, cost []int
		expected  int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}, 3},
		{[]int{2, 3, 4}, []int{3, 4, 3}, -1},
	}
	for _, v := range e {
		o := canCompleteCircuit(v.gas, v.cost)
		if o != v.expected {
			t.Error(fmt.Sprintf("%v %v", v, o))
		}
	}
}
