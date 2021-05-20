package Num_221_MaximalSquare

import (
	"fmt"
	"testing"
)

func TestNum_221_MaximalSquare(t *testing.T) {

	e := []struct {
		matrix   [][]byte
		expected int
	}{
		{matrix: [][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}}, expected: 4},
		{matrix: [][]byte{{'0', '1'}, {'1', '0'}}, expected: 1},
		{matrix: [][]byte{{'1'}}, expected: 1},
		{matrix: [][]byte{{'0'}}, expected: 0},
	}
	for _, v := range e {
		o := maximalSquare(v.matrix)
		if o != v.expected {
			t.Error(fmt.Sprintf("%v %v", v, o))
		}
	}
}
