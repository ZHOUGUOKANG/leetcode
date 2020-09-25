package Num_78_Subsets

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNum_78_Subsets(t *testing.T) {
	e := []struct {
		nums     []int
		expected [][]int
	}{
		{[]int{1, 2}, [][]int{{1}, {1, 2}, {2}, {}}},
		{[]int{1, 2, 3}, [][]int{{1, 2}, {1, 2, 3}, {1, 3}, {1}, {2}, {2, 3}, {3}, {}}},
	}
	for _, v := range e {
		o := subsets(v.nums)
		if !reflect.DeepEqual(v.expected, o) {
			t.Error(fmt.Sprintf("%v %v", v.expected, o))
		}
	}
}
