package Num_32_longestValidParentheses

import "testing"

func TestLongestValidParentheses(t *testing.T) {
	in := "(()"
	expected := 2
	out := longestValidParentheses(in)
	if expected != out {
		t.Error(in, expected, out)
	}
}

func TestLongestValidParentheses1(t *testing.T) {
	in := ")()())"
	expected := 4
	out := longestValidParentheses(in)
	if expected != out {
		t.Error(in, expected, out)
	}
}
