package Num_offer_08_01_waysToStep

import "testing"

func TestWaysToStep(t *testing.T) {
	input := 3
	expected := 4
	output := waysToStep(input)
	if output != expected {
		t.Error(input, expected, output)
	}
}

func TestWaysToStep1(t *testing.T) {
	input := 5
	expected := 13
	output := waysToStep(input)
	if output != expected {
		t.Error(input, expected, output)
	}
}
