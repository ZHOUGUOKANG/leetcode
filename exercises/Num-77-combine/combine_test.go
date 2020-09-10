package Num_77_combine

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCombine(t *testing.T) {
	example := []struct {
		n        int
		k        int
		expected [][]int
	}{
		{4, 2, [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}}},
		//{2,1,[][]int{{1},{2}}},
		//{4,3,[][]int{{1,2,3},{1,2,4},{1,3,4},{2,3,4}}},
	}
	for k, v := range example {
		o := combine(v.n, v.k)
		if !reflect.DeepEqual(o, v.expected) {
			t.Error(k, fmt.Sprintf("%v %v", v, o))
		}
	}
}
