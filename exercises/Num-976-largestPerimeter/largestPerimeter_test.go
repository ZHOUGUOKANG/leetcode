package Num_976_LargestPerimeter

import (
	"fmt"
	"testing"
)

func TestNum_976_LargestPerimeter(t *testing.T) {

	e := []struct {
		A        []int
		expected int
	}{
		{[]int{2, 1, 2}, 5},
	}
	for _, v := range e {
		o := largestPerimeter(v.A)
		if o != v.expected {
			t.Error(fmt.Sprintf("%v %v", v, o))
		}
	}
}
