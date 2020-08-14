package Num_29_divide

import "testing"

func TestDivide(t *testing.T) {
	dividend, divisor := 10, 3
	expected := 3
	out := divide(dividend, divisor)
	if expected != out {
		t.Error(dividend, divisor, expected, out)
	}
}

func TestDivide1(t *testing.T) {
	dividend, divisor := 7, -3
	expected := -2
	out := divide(dividend, divisor)
	if expected != out {
		t.Error(dividend, divisor, expected, out)
	}
}
