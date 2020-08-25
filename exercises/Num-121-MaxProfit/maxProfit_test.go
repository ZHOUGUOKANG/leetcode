package Num_121_MaxProfit

import "testing"

func TestMaxProfit(t *testing.T) {
	input := []int{7, 1, 5, 3, 6, 4}
	expected := 5
	output := MaxProfit(input)
	if output != expected {
		t.Error("expected", expected, "output", output)
	}

	input = []int{4, 2, 3, 3, 1, 6}
	expected = 5
	output = MaxProfit(input)
	if output != expected {
		t.Error("expected", expected, "output", output)
	}

	input = []int{7, 1, 7, 3, 1, 6}
	expected = 6
	output = MaxProfit(input)
	if output != expected {
		t.Error("expected", expected, "output", output)
	}
}
