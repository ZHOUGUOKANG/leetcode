package Num_1289_minFallingPathSum

import "testing"

func TestMinFallingPathSum(t *testing.T) {
	input := [][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}}
	output := minFallingPathSum(input)
	if output != 13 {
		t.Error(input, 13, output)
	}

	input = [][]int{[]int{1, 2, 3}, []int{2, 3, 1}, []int{6, 8, 9}}
	output = minFallingPathSum(input)
	if output != 8 {
		t.Error(input, 8, output)
	}
}
