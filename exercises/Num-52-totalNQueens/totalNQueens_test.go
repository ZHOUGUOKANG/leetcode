package Num_52_totalNQueens

import "testing"

func TestNum_52_totalNQueens(t *testing.T) {
	examples := []struct {
		n, expected int
	}{
		{1, 1},
		{2, 0},
		{3, 0},
		{4, 2},
		{6, 4},
		{9, 352},
	}
	for k, v := range examples {
		o := totalNQueens(v.n)
		if o != v.expected {
			t.Error(k, v.n, v.expected, o)
		}
	}
}
