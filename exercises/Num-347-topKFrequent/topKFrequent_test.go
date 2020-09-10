package Num_347_topKFrequent

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	examples := []struct {
		nums     []int
		k        int
		expected []int
	}{
		{[]int{1, 1, 1, 2, 2, 3}, 2, []int{1, 2}},
		{[]int{1}, 1, []int{1}},
		{[]int{1, 1, 1, 1, 1, 1, 3, 3, 3, 3, 3, 0, 2, 2, 2, 2}, 2, []int{1, 3}},
	}

	for k, v := range examples {
		o := topKFrequent(v.nums, v.k)
		if !reflect.DeepEqual(o, v.expected) {
			t.Error(k, fmt.Sprintf("%v %v", v, o))
		}
	}
}
