package Num_374_GuessNumber

import (
	"fmt"
	"testing"
)

func TestNum_374_GuessNumber(t *testing.T) {

	e := []struct {
		input    int
		expected int
	}{
		{input: 10, expected: 6},
		{input: 100, expected: 66},
		{input: 150, expected: 32},
		{input: 32543, expected: 66},
	}
	for _, v := range e {
		guessNum = v.expected
		o := guessNumber(v.input)
		//t.Error(fmt.Sprintf("%v %v", v, o))
		if o != v.expected {
			t.Error(fmt.Sprintf("%v %v", v, o))
		}
	}
}
