package Num_718_findLength

import "testing"

func TestFindLength(t *testing.T) {
	input1, input2 := []int{1, 2, 3, 2, 1}, []int{3, 2, 1, 4, 7}
	expected := 3
	output := findLength(input1, input2)
	if output != expected {
		t.Error(input1, input2, expected, output)
	}
}
