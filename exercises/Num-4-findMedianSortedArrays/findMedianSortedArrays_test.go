package Num_4_findMedianSortedArrays

import "testing"

func TestFindMedianSortedArrays(t *testing.T) {
	input1 := []int{1, 3}
	input2 := []int{2}
	expected := float64(2)
	output := findMedianSortedArrays(input1, input2)
	if output != expected {
		t.Error(input1, input2, expected, output)
	}

	input1 = []int{1, 2}
	input2 = []int{3, 4}
	expected = 2.5
	output = findMedianSortedArrays(input1, input2)
	if output != expected {
		t.Error(input1, input2, expected, output)
	}

	input1 = []int{1, 3, 4, 9}
	input2 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected = 4.5
	output = findMedianSortedArrays(input1, input2)
	if output != expected {
		t.Error(input1, input2, expected, output)
	}
}
