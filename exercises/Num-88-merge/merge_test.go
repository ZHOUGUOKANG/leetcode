package Num_88_Merge

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNum_88_Merge(t *testing.T) {
	e := []struct {
		num1, num2 []int
		m, n       int
		expected   []int
	}{
		{[]int{1, 2, 3, 0, 0, 0}, []int{2, 5, 6}, 3, 3, []int{1, 2, 2, 3, 5, 6}},
		{[]int{4, 5, 6, 0, 0, 0}, []int{1, 2, 3}, 3, 3, []int{1, 2, 3, 4, 5, 6}},
		{[]int{4, 0, 0, 0, 0, 0}, []int{1, 2, 3, 5, 6}, 1, 5, []int{1, 2, 3, 4, 5, 6}},
	}
	for _, v := range e {
		merge(v.num1, v.m, v.num2, v.n)
		if !reflect.DeepEqual(v.num1, v.expected) {
			t.Error(fmt.Sprintf("%v", v))
		}
	}
}
