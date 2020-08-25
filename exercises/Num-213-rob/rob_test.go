package Num_213_rob

import "testing"

func TestRob(t *testing.T) {
	nums := []int{2, 3, 2}
	expected := 3
	out := rob(nums)
	if expected != out {
		t.Error(nums, expected, out)
	}
}

func TestRob1(t *testing.T) {
	nums := []int{1, 2, 3, 1}
	expected := 4
	out := rob(nums)
	if expected != out {
		t.Error(nums, expected, out)
	}
}

func TestRob2(t *testing.T) {
	nums := []int{2, 1, 1, 2}
	expected := 3
	out := rob(nums)
	if expected != out {
		t.Error(nums, expected, out)
	}
}
