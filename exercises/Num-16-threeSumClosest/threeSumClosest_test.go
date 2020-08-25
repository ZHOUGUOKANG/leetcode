package Num_16_threeSumClosest

import "testing"

func TestThreeSum(t *testing.T) {
	in, target := []int{-1, 2, 1, -4}, 1
	expected := 2
	out := threeSumClosest(in, target)
	if expected != out {
		t.Error(in, target, expected, out)
	}
}
