package Num_28_strStr

import "testing"

type test struct {
	haystack string
	needle   string
	expected int
}

func TestStrStr(t *testing.T) {
	tests := []*test{
		{"hello", "ll", 2},
		{"aaaaa", "bba", -1},
		{"aaa", "aaaa", -1},
		{"mississippi", "mississippi", 0},
	}
	for _, v := range tests {
		if out := strStr(v.haystack, v.needle); out != v.expected {
			t.Error(v, out)
		}
	}
}
