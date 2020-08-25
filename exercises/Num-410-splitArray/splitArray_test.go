package Num_410_splitArray

import "testing"

func TestSplitArray(t *testing.T) {
	input1 := []int{7, 2, 5, 10, 8}
	input2 := 2
	expected := 18
	out := splitArray(input1, input2)
	if out != expected {
		t.Error(input1, input2, expected, out)
	}
}
