package Num_17_16_massage

import "testing"

func TestMassage(t *testing.T) {
	input := []int{1, 2, 3, 1}
	expected := 4
	output := massage(input)
	if output != expected {
		t.Error(input, expected, output)
	}
}

func TestMassage1(t *testing.T) {
	input := []int{2, 7, 9, 3, 1}
	expected := 12
	output := massage(input)
	if output != expected {
		t.Error(input, expected, output)
	}
}

func TestMassage2(t *testing.T) {
	input := []int{2, 1, 4, 5, 3, 1, 1, 3}
	expected := 12
	output := massage(input)
	if output != expected {
		t.Error(input, expected, output)
	}
}
