package Num_1021_removeOuterParentheses

import "testing"

func TestRemoveOuterParentheses(t *testing.T) {
	var in, expected, output string = "(()())(())", "()()()", ""
	output = RemoveOuterParentheses(in)
	if output != expected {
		t.Error("expected", expected, "output", output)
	}
	in, expected = "(()())(())(()(()))", "()()()()(())"
	output = RemoveOuterParentheses(in)
	if output != expected {
		t.Error("expected", expected, "output", output)
	}
	in, expected = "(()())(())(()(()))", "()()()()(())"
	output = RemoveOuterParentheses(in)
	if output != expected {
		t.Error("expected", expected, "output", output)
	}
	in, expected = "()()", ""
	output = RemoveOuterParentheses(in)
	if output != expected {
		t.Error("expected", expected, "output", output)
	}
}
