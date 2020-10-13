package Num_75_SortColors

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNum_75_SortColors(t *testing.T) {
	e := []struct {
		nums     []int
		expected []int
	}{
		{[]int{2, 0, 2, 1, 1, 0}, []int{0, 0, 1, 1, 2, 2}},
	}
	for _, v := range e {
		sortColors(v.nums)
		if !reflect.DeepEqual(v.nums, v.expected) {
			t.Error(fmt.Sprintf("%v", v))
		}
	}
}
