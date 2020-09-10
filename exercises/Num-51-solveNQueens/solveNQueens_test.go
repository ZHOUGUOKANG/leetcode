package Num_51_solveNQueens

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSolveNQueens(t *testing.T) {
	examples := []struct {
		n        int
		expected [][]string
	}{
		{1, [][]string{{"Q"}}},
		{2, [][]string{}},
		{3, [][]string{}},
		{4, [][]string{{".Q..", "...Q", "Q...", "..Q."}, {"..Q.", "Q...", "...Q", ".Q.."}}},
		{6, [][]string{{".Q....", "...Q..", ".....Q", "Q.....", "..Q...", "....Q."}, {"..Q...", ".....Q", ".Q....", "....Q.", "Q.....", "...Q.."}, {"...Q..", "Q.....", "....Q.", ".Q....", ".....Q", "..Q..."}, {"....Q.", "..Q...", "Q.....", ".....Q", "...Q..", ".Q...."}}},
	}

	for k, v := range examples {
		o := solveNQueens(v.n)
		if !reflect.DeepEqual(o, v.expected) {
			t.Error(k, fmt.Sprintf("%v %v", o, v))
		}
	}
}
