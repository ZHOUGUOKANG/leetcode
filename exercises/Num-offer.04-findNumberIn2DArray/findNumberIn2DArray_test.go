package Num_offer_04_findNumberIn2DArray

import (
	"fmt"
	"testing"
)

func TestFind(t *testing.T) {
	examples := []struct {
		nums     [][]int
		target   int
		expected bool
	}{
		{[][]int{{-5}}, -5, true},
		{[][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, 5, true},
		{[][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, 20, false},
	}

	for k, v := range examples {
		if findNumberIn2DArray(v.nums, v.target) != v.expected {
			t.Error(k, fmt.Sprintf("%v", v))
		}
	}
}
