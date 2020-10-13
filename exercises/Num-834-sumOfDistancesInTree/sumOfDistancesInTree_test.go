package Num_834_SumOfDistancesInTree

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNum_834_SumOfDistancesInTree(t *testing.T) {
	e := []struct {
		nums     [][]int
		n        int
		expected []int
	}{
		{[][]int{{0, 1}, {0, 2}, {2, 3}, {2, 4}, {2, 5}}, 6, []int{8, 12, 6, 10, 10, 10}},
	}
	for _, v := range e {
		o := sumOfDistancesInTree(v.n, v.nums)
		if !reflect.DeepEqual(v.expected, o) {
			t.Error(fmt.Sprintf("%v %v", v, o))
		}
	}
}
