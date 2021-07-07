package Num_645_FindErrorNums

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNum_645_FindErrorNums(t *testing.T) {
	e := []struct {
		input    []int
		expected []int
	}{
		{input: []int{1, 2, 2, 4}, expected: []int{2, 3}},
		{input: []int{1, 1}, expected: []int{1, 2}},
		{input: []int{1, 2, 3, 5, 5, 4, 7, 8}, expected: []int{5, 6}},
	}
	for _, v := range e {
		o := findErrorNums(v.input)
		if !reflect.DeepEqual(v.expected, o) {
			t.Error(fmt.Sprintf("%v %v", v, o))
		}
	}
}
