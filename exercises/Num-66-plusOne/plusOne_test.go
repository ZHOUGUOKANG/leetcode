package Num_66_PlusOne

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNum_66_PlusOne(t *testing.T) {
	e := []struct {
		nums     []int
		expected []int
	}{
		{[]int{1, 2, 3}, []int{1, 2, 4}},
		{[]int{4, 3, 2, 1}, []int{4, 3, 2, 2}},
		{[]int{9, 9, 9}, []int{1, 0, 0, 0}},
		{[]int{1, 9, 9}, []int{2, 0, 0}},
		{[]int{0}, []int{1}},
	}
	for _, v := range e {
		o := plusOne(v.nums)
		if !reflect.DeepEqual(o, v.expected) {
			t.Error(fmt.Sprintf("%v%v", v, o))
		}
	}
}
