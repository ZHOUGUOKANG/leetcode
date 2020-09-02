package Num_48_rotate

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	tests := []struct {
		input    [][]int
		expected [][]int
	}{
		{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}}},
		{[][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}, [][]int{{15, 13, 2, 5}, {14, 3, 4, 1}, {12, 6, 8, 9}, {16, 7, 10, 11}}},
	}

	for _, v := range tests {
		var in [][]int
		copy(in, v.input)
		rotate(v.input)
		if !reflect.DeepEqual(v.input, v.expected) {
			t.Error(fmt.Sprintf("%v %v  %v", in, v.expected, v.input))
		}
	}
}
