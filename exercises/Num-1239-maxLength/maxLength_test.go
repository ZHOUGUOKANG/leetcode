package Num_1239_MaxLength

import (
	"fmt"
	"testing"
)

func TestNum_1239_MaxLength(t *testing.T) {
	e := []struct {
		input    []string
		expected int
	}{
		{input: []string{"un", "iq", "ue"}, expected: 4},
		{input: []string{"cha", "r", "act", "ers"}, expected: 6},
		{input: []string{"abcdefghijklmnopqrstuvwxyz"}, expected: 26},
	}
	for _, v := range e {
		o := maxLength(v.input)
		if o != v.expected {
			t.Error(fmt.Sprintf("%v %v", v, o))
		}
	}
}
