package Num_56_Merge

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNum_56_Merge(t *testing.T) {
	example := []struct {
		nums     [][]int
		expected [][]int
	}{
		{[][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, [][]int{{1, 6}, {8, 10}, {15, 18}}},
		{[][]int{{1, 4}, {4, 5}}, [][]int{{1, 5}}},
		{[][]int{{2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10}}, [][]int{{1, 10}}},
	}
	for _, v := range example {
		o := merge(v.nums)
		if !reflect.DeepEqual(o, v.expected) {
			t.Error(fmt.Sprintf("%v %v", v, o))
		}
	}
}
