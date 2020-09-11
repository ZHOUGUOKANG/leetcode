package Num_54_spiralOrder

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNum_54_spiralOrder(t *testing.T) {
	examples := []struct {
		nums     [][]int
		expected []int
	}{
		{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, []int{1, 2, 3, 6, 9, 8, 7, 4, 5}},
		{[][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}, []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}},
	}

	for _, v := range examples {
		o := spiralOrder(v.nums)
		if !reflect.DeepEqual(o, v.expected) {
			t.Error(fmt.Sprintf("%v", v), o)
		}
	}
}
