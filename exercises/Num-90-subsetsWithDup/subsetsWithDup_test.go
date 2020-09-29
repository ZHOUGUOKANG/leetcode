package Num_90_SubsetsWithDup

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNum_90_SubsetsWithDup(t *testing.T) {
	e := []struct {
		nums     []int
		expected [][]int
	}{
		{[]int{1, 2, 2}, [][]int{{}, {1}, {1, 2}, {1, 2, 2}, {2}, {2, 2}}},
	}
	for _, v := range e {
		o := subsetsWithDup(v.nums)
		if !reflect.DeepEqual(v.expected, o) {
			t.Error(fmt.Sprintf("%v %v", v, o))
		}
	}
}
