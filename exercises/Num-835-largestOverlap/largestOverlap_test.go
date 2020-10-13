package Num_835_LargestOverlap

import (
	"fmt"
	"testing"
)

func TestNum_835_LargestOverlap(t *testing.T) {
	e := []struct {
		nums1    [][]int
		nums2    [][]int
		expected int
	}{
		{[][]int{{1, 1, 0}, {0, 1, 0}, {0, 1, 0}}, [][]int{{0, 0, 0}, {0, 1, 1}, {0, 0, 1}}, 3},
		{[][]int{{0, 1}, {1, 1}}, [][]int{{1, 1}, {1, 0}}, 2},
	}
	for _, v := range e {
		o := largestOverlap(v.nums1, v.nums2)
		if o != v.expected {
			t.Error(fmt.Sprintf("%v %v", v, o))
		}
	}
}
