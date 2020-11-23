package Num_452_FindMinArrowShots

import (
	"fmt"
	"testing"
)

func TestNum_452_FindMinArrowShots(t *testing.T) {

	e := []struct {
		input    [][]int
		expected int
	}{
		{[][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}, 2},
		{[][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}, 4},
		{[][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}, 2},
		{[][]int{{1, 2}}, 1},
		{[][]int{{2, 3}, {2, 3}}, 1},
	}
	for _, v := range e {
		o := findMinArrowShots(v.input)
		if o != v.expected {
			t.Error(fmt.Sprintf("%v %v", v, o))
		}
	}
}
