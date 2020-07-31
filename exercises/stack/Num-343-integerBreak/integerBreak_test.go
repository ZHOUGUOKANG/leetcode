package Num_343_integerBreak

import "testing"

func TestIntegerBreak(t *testing.T) {
	in := 2
	expected := 1
	out := integerBreak(in)
	if out != expected {
		t.Error(in, expected, out)
	}
}
func TestIntegerBreak1(t *testing.T) {
	in := 10
	expected := 36
	out := integerBreak(in)
	if out != expected {
		t.Error(in, expected, out)
	}
}
func TestIntegerBreak2(t *testing.T) {
	in := 3
	expected := 2
	out := integerBreak(in)
	if out != expected {
		t.Error(in, expected, out)
	}
}
