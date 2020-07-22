package Num_offer_42_maxSubArray

import "testing"

func TestMaxSubArray(t *testing.T) {
	input := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	expected := 6
	output := maxSubArray(input)
	if expected != output {
		t.Error(input, expected, output)
	}
}

func TestMaxSubArray1(t *testing.T) {
	input := []int{1}
	expected := 1
	output := maxSubArray(input)
	if expected != output {
		t.Error(input, expected, output)
	}
}

func TestMaxSubArray2(t *testing.T) {
	input := []int{-1}
	expected := -1
	output := maxSubArray(input)
	if expected != output {
		t.Error(input, expected, output)
	}
}
