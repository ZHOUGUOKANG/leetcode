package Num_16_17_maxSubArray

import "testing"

func TestMaxSubArray(t *testing.T) {
	input := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	output := maxSubArray(input)
	if output != 6 {
		t.Error(input, output, 6)
	}
}
