package Num_283_MoveZeroes

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNum_283_MoveZeroes(t *testing.T) {
	e := []struct {
		nums, expected []int
	}{
		{[]int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
	}
	for _, v := range e {
		moveZeroes(v.nums)
		if !reflect.DeepEqual(v.expected, v.nums) {
			t.Error(fmt.Sprintf("%v %v", v, v.nums))
		}
	}
}
