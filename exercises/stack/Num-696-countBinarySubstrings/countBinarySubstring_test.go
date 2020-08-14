package Num_696_countBinarySubstrings

import "testing"

func TestCountBinarySubstrings(t *testing.T) {
	s := "00110011"
	expected := 6
	out := countBinarySubstrings(s)
	if expected != out {
		t.Error(s, expected, out)
	}
}
func TestCountBinarySubstrings1(t *testing.T) {
	s := "10101"
	expected := 4
	out := countBinarySubstrings(s)
	if expected != out {
		t.Error(s, expected, out)
	}
}
