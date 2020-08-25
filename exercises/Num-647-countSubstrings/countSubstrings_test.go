package Num_647_countSubstrings

import "testing"

func TestCountSubstrings(t *testing.T) {
	tests := []*struct {
		s        string
		expected int
	}{
		{"abc", 3},
		{"aaa", 6},
		{"a", 1},
		{"aa", 3},
		{"aaaaa", 15},
		{"krggyuxvmoobsdchrgdwlopeykbdjzlln", 36},
	}
	for _, v := range tests {
		out := countSubstrings(v.s)
		if out != v.expected {
			t.Error(v, out)
		}
	}
}
