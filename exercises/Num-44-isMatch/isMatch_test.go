package Num_44_isMatch

import "testing"

func TestIsMatch(t *testing.T) {
	in1, in2 := "aa", "a"
	expected := false
	out := isMatch(in1, in2)
	if out != expected {
		t.Error(in1, in2, expected, out)
	}
}
func TestIsMatch1(t *testing.T) {
	in1, in2 := "aa", "*"
	expected := true
	out := isMatch(in1, in2)
	if out != expected {
		t.Error(in1, in2, expected, out)
	}
}
func TestIsMatch2(t *testing.T) {
	in1, in2 := "cb", "*a"
	expected := false
	out := isMatch(in1, in2)
	if out != expected {
		t.Error(in1, in2, expected, out)
	}
}
func TestIsMatch3(t *testing.T) {
	in1, in2 := "adceb", "*a*b"
	expected := true
	out := isMatch(in1, in2)
	if out != expected {
		t.Error(in1, in2, expected, out)
	}
}
func TestIsMatch4(t *testing.T) {
	in1, in2 := "acdcb", "a*c?b"
	expected := false
	out := isMatch(in1, in2)
	if out != expected {
		t.Error(in1, in2, expected, out)
	}
}
