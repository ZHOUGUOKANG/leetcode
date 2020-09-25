package Num_74_SearchMatrix

import (
	"fmt"
	"testing"
)

func TestNum_74_SearchMatrix(t *testing.T) {
	e := []struct {
		matrix   [][]int
		target   int
		expected bool
	}{
		{[][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 50}}, 3, true},
		{[][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 50}}, 13, false},
		{[][]int{{1}}, 0, false},
		{[][]int{{1}}, 1, true},
		{[][]int{{1}, {3}}, 3, true},
	}
	for _, v := range e {
		o := searchMatrix(v.matrix, v.target)
		if o != v.expected {
			t.Error(fmt.Sprintf("%v %v", v, o))
		}
	}
}
