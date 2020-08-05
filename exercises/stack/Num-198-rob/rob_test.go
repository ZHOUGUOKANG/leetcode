package Num_198_rob

import "testing"

func TestRob(t *testing.T) {
	nums := []int{1, 2, 3, 1}
	expected := 4
	out := rob(nums)
	if expected != out {
		t.Error(nums, expected, out)
	}
}

func TestRob1(t *testing.T) {
	nums := []int{2, 7, 9, 3, 1}
	expected := 12
	out := rob(nums)
	if expected != out {
		t.Error(nums, expected, out)
	}
}
