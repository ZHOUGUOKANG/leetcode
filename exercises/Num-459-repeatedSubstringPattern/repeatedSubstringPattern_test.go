package Num_459_repeatedSubstringPattern

import "testing"

func TestRepeatedSubstringPattern(t *testing.T) {
	tests := []*struct {
		s        string
		excepted bool
	}{
		{"abab", true},
		{"abcabcabcabc", true},
		{"aba", false},
		{"aabaaba", false},
		{"abaababaab", true},
	}

	for _, v := range tests {
		out := repeatedSubstringPattern(v.s)
		if v.excepted != out {
			t.Error(v)
		}
	}
}
