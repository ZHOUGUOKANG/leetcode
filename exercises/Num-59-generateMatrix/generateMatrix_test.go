package Num_59_GenerateMatrix

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNum_59_GenerateMatrix(t *testing.T) {
	examples := []struct {
		n        int
		expected [][]int
	}{
		{3, [][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}}},
	}

	for _, v := range examples {
		o := generateMatrix(v.n)
		if !reflect.DeepEqual(o, v.expected) {
			t.Error(v.n, fmt.Sprintf("%v%v", v.expected, o))
		}
	}
}
