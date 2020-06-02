package Num_1218_longestSubsequence

import "testing"

func TestLongestSubsequence(t *testing.T) {
	input1, input2 := []int{1, 2, 3, 4}, 1
	output := longestSubsequence(input1, input2)
	if output != 4 {
		t.Error(input1, input2, output)
	}

	input1, input2 = []int{1, 3, 5, 7}, 1
	output = longestSubsequence(input1, input2)
	if output != 1 {
		t.Error(input1, input2, output)
	}

	input1, input2 = []int{1, 5, 7, 8, 5, 3, 4, 2, 1}, -2
	output = longestSubsequence(input1, input2)
	if output != 4 {
		t.Error(input1, input2, output)
	}
}
