package Num_16_05_trailingZeroes

import "testing"

func TestTrailingZeroes(t *testing.T) {
	in := 3
	expected := 0
	out := trailingZeroes(in)
	if expected != out {
		t.Error(in, expected, out)
	}
}

func TestTrailingZeroes1(t *testing.T) {
	in := 5
	expected := 1
	out := trailingZeroes(in)
	if expected != out {
		t.Error(in, expected, out)
	}
}
