package Num_832_FlipAndInvertImage

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNum_832_FlipAndInvertImage(t *testing.T) {
	e := []struct {
		nums     [][]int
		expected [][]int
	}{
		{[][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}}, [][]int{{1, 0, 0}, {0, 1, 0}, {1, 1, 1}}},
		{[][]int{{1, 1, 0, 0}, {1, 0, 1, 0}}, [][]int{{1, 1, 0, 0}, {1, 0, 1, 0}}},
	}
	for _, v := range e {
		o := flipAndInvertImage(v.nums)
		if !reflect.DeepEqual(v.expected, o) {
			t.Error(fmt.Sprintf("%v %v", v, o))
		}
	}
}
