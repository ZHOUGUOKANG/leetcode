package Num_64_minPathSum

import "testing"

func TestMinPathSum(t *testing.T) {
	input := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	expected := 7
	output := minPathSum(input)
	if output != expected {
		t.Error(input, expected, output)
	}
}
