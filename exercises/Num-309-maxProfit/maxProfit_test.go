package Num_309_maxProfit

import "testing"

func TestMaxProfit(t *testing.T) {
	expected := 3
	in := []int{1, 2, 3, 0, 2}
	out := maxProfit(in)
	if expected != out {
		t.Error(in, expected, out)
	}
}
