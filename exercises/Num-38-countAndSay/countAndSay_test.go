package Num_38_countAndSay

import "testing"

func TestCountAndSay(t *testing.T) {
	tests := []*struct {
		n        int
		expected string
	}{
		{1, "1"},
		{2, "11"},
		{3, "21"},
		{4, "1211"},
		{5, "111221"},
	}
	for _, v := range tests {
		out := countAndSay(v.n)
		if out != v.expected {
			t.Error(v, out)
		}
	}
}
