package Num_5_LongestPalindrome

import "testing"

func TestLongestPalindrome(t *testing.T) {
	in := "babad"
	out := LongestPalindrome(in)
	expected1 := "abd"
	expected2 := "bab"
	if !(out == expected1 || out == expected2) {
		t.Error(in, out, expected1, expected2)
	}
}
