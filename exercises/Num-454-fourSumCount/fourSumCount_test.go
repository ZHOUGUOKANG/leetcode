package Num_454_FourSumCount

import (
	"fmt"
	"testing"
)

func TestNum_454_FourSumCount(t *testing.T) {

	e := []struct {
		A, B, C, D []int
		expected   int
	}{
		{[]int{1, 2}, []int{-2, -1}, []int{-1, 2}, []int{0, 2}, 2},
	}
	for _, v := range e {
		o := fourSumCount(v.A, v.B, v.C, v.D)
		if o != v.expected {
			t.Error(fmt.Sprintf("%v %v", v, o))
		}
	}
}
