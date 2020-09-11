package Num_216_combinationSum3

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNum_216_combinationSum3(t *testing.T) {
	examples := []struct {
		k, n     int
		expected [][]int
	}{
		{3, 7, [][]int{{1, 2, 4}}},
		{3, 9, [][]int{{1, 2, 6}, {1, 3, 5}, {2, 3, 4}}},
		{2, 18, nil},
	}

	for k, v := range examples {
		o := combinationSum3(v.k, v.n)
		if !reflect.DeepEqual(o, v.expected) {
			t.Error(k, fmt.Sprintf("%v %v", v, o))
		}
	}
}
