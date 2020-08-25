package Num_174_calculateMinimumHP

import "testing"

func TestCalculateMinimumHP(t *testing.T) {
	output := calculateMinimumHP([][]int{{-2, -3, 3}, {-5, -10, 1}, {10, 30, -5}})
	expected := 7
	if output != expected {
		t.Error(output, expected)
	}

	output = calculateMinimumHP([][]int{{2}, {1}})
	expected = 1
	if output != expected {
		t.Error(output, expected)
	}

	output = calculateMinimumHP([][]int{{0, 0, 0}, {1, 1, -1}})
	expected = 1
	if output != expected {
		t.Error(output, expected)
	}

	output = calculateMinimumHP([][]int{{0, -3}})
	expected = 4
	if output != expected {
		t.Error(output, expected)
	}
}
