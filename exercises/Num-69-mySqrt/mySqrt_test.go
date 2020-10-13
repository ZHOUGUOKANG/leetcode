package Num_69_MySqrt

import (
	"fmt"
	"testing"
)

func TestNum_69_MySqrt(t *testing.T) {
	e := []struct {
		num      int
		expected int
	}{
		{0, 0},
		{1, 1},
		{4, 2},
		{8, 2},
		{9, 3},
		{12, 3},
		{17, 4},
	}
	for _, v := range e {
		o := mySqrt(v.num)
		if o != v.expected {
			t.Error(fmt.Sprintf("%v %v", v, o))
		}
	}
}
