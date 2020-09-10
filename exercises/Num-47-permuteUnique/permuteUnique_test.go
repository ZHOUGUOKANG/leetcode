package Num_47_permuteUnique

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNum_47_permuteUnique(t *testing.T) {
	examples := []struct {
		nums     []int
		expected [][]int
	}{
		{[]int{1, 1, 2}, [][]int{{1, 1, 2}, {1, 2, 1}, {2, 1, 1}}},
		{[]int{3, 3, 0, 3}, [][]int{{0, 3, 3, 3}, {3, 0, 3, 3}, {3, 3, 0, 3}, {3, 3, 3, 0}}},
	}

	for k, v := range examples {
		o := permuteUnique(v.nums)
		if !reflect.DeepEqual(o, v.expected) {
			t.Error(k, fmt.Sprintf("%v %v", v, o))
		}
	}
}
