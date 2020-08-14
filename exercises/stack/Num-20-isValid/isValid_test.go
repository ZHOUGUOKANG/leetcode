package Num_20_isValid

import "testing"

type testIsValid struct {
	input    string
	expected bool
}

func TestIsValid(t *testing.T) {
	tests := []*testIsValid{
		{"()", true},
		{"(", false},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"({[[]]})", true},
	}
	for _, v := range tests {
		output := isValid(v.input)
		if output != v.expected {
			t.Error(v, output)
		}
	}
}
