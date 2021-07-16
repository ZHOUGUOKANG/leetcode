package Num_53_Search

import (
	"fmt"
	"testing"
)

func TestNum_53_Search(t *testing.T) {
	e := []struct {
		nums     []int
		target   int
		expected int
	}{
		{nums: []int{5, 7, 7, 8, 8, 10}, target: 8, expected: 2},
		{nums: []int{5, 7, 7, 8, 8, 10}, target: 5, expected: 1},
		{nums: []int{5, 7, 7, 8, 8, 10}, target: 7, expected: 2},
		{nums: []int{5, 7, 7, 8, 8, 10}, target: 10, expected: 1},
		{nums: []int{5, 7, 7, 8, 8, 10}, target: 0, expected: 0},
		{nums: []int{0}, target: 0, expected: 0},
		{nums: []int{}, target: 0, expected: 0},
	}
	for _, v := range e {
		o := search(v.nums, v.target)
		if o != v.expected {
			t.Error(fmt.Sprintf("%v %v", v, o))
		}
	}
}
