package Num_offer_11_minArray

import "testing"

func TestMinArray(t *testing.T) {
	input := []int{3, 4, 5, 1, 2}
	expected := 1
	output := minArray(input)
	if output != expected {
		t.Error(input, expected, output)
	}
}
func TestMinArray1(t *testing.T) {
	input := []int{2, 2, 2, 0, 1}
	expected := 0
	output := minArray(input)
	if output != expected {
		t.Error(input, expected, output)
	}
}

func TestMinArray2(t *testing.T) {
	input := []int{1, 3, 5}
	expected := 1
	output := minArray(input)
	if output != expected {
		t.Error(input, expected, output)
	}
}
