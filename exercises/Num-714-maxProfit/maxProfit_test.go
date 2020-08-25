package Num_714_maxProfit

import "testing"

func TestMaxProfit(t *testing.T) {
	input1 := []int{1, 3, 2, 8, 4, 9}
	input2 := 2
	expected := 8
	output := maxProfit(input1, input2)
	if expected != output {
		t.Error(input1, input2, expected, output)
	}
}
