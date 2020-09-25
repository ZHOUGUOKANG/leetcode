package Num_73_SetZeroes

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNum_73_SetZeroes(t *testing.T) {
	e := []struct {
		nums     [][]int
		expected [][]int
	}{
		{[][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}, [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}},
		{[][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}, [][]int{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}}},
		{[][]int{{1, 2, 3, 4}, {5, 0, 7, 8}, {0, 10, 11, 12}, {13, 14, 15, 0}}, [][]int{{0, 0, 3, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}}},
	}
	for _, v := range e {
		setZeroes(v.nums)
		if !reflect.DeepEqual(v.nums, v.expected) {
			t.Error(fmt.Sprintf("%v", v))
		}
	}
}
