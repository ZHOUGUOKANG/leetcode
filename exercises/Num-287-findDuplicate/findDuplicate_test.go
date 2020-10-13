package Num_287_FindDuplicate

import (
	"fmt"
	"testing"
)

func TestNum_287_FindDuplicate(t *testing.T) {

	e := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 3, 4, 2, 2}, 2},
		{[]int{3, 1, 3, 4, 2}, 3},
	}
	for _, v := range e {
		o := findDuplicate(v.nums)
		if o != v.expected {
			t.Error(fmt.Sprintf("%v %v", v, o))
		}
	}
}
