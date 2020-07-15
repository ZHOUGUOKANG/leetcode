package Num_120_minimumTotal

import "testing"

func TestMinimumTotal(t *testing.T) {
	expected := 11
	input := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	output := minimumTotal(input)
	if expected != output {
		t.Error(output, expected)
	}
}
