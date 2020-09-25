package Num_70_ClimbStairs

import (
	"fmt"
	"testing"
)

func TestNum_70_ClimbStairs(t *testing.T) {
	e := []struct {
		n        int
		expected int
	}{
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 5},
		{5, 8},
	}
	for _, v := range e {
		o := climbStairs(v.n)
		if o != v.expected {
			t.Error(fmt.Sprintf("%v %v", v, o))
		}
	}
}
