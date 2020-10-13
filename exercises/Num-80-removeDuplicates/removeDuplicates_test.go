package Num_80_RemoveDuplicates

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNum_80_RemoveDuplicates(t *testing.T) {
	e := []struct {
		nums         []int
		expected     int
		expectedNums []int
	}{
		{[]int{1, 1, 1, 2, 2, 3}, 5, []int{1, 1, 2, 2, 3}},
		{[]int{0, 0, 1, 1, 1, 1, 2, 3, 3}, 7, []int{0, 0, 1, 1, 2, 3, 3}},
		{[]int{0, 0}, 2, []int{0, 0}},
		{[]int{0}, 1, []int{0}},
	}
	for _, v := range e {
		o := removeDuplicates(v.nums)
		if o != v.expected || !reflect.DeepEqual(v.nums[:v.expected], v.expectedNums) {
			t.Error(fmt.Sprintf("%v %v", v, o))
		}
	}
}
