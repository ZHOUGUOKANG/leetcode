package Num_279_NumSquares

import (
	"fmt"
	"testing"
)

func TestNum_279_NumSquares(t *testing.T) {
	e := []struct {
		n        int
		expected int
	}{
		{n: 1, expected: 1},
		{n: 2, expected: 2},
		{n: 3, expected: 3},
		{n: 4, expected: 1},
		{n: 5, expected: 2},
		{n: 10, expected: 2},
		{n: 12, expected: 3},
		{n: 13, expected: 2},
	}
	for _, v := range e {
		o := numSquares(v.n)
		if o != v.expected {
			t.Error(fmt.Sprintf("%v %v", v, o))
		}
	}
}
