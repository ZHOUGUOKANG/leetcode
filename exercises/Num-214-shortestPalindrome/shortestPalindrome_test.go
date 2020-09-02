package Num_214_shortestPalindrome

import "testing"

func TestShortestPalindrome(t *testing.T) {
	examples := []*struct {
		s        string
		expected string
	}{
		{"aacecaaa", "aaacecaaa"},
		{"abcd", "dcbabcd"},
		{"", ""},
		{"abbacd", "dcabbacd"},
		{"a", "a"},
		{"abb", "bbabb"},
		{"aaaabbaa", "aabbaaaabbaa"},
	}

	for _, v := range examples {
		o := shortestPalindrome(v.s)
		if o != v.expected {
			t.Error(v.s, v.expected, o)
		}
	}
}
