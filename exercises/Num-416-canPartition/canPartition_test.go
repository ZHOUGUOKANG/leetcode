package Num_416_CanPartition

import (
	"fmt"
	"testing"
)

func TestNum_416_CanPartition(t *testing.T) {

	e := []struct {
		nums     []int
		expected bool
	}{
		{[]int{1, 5, 11, 5}, true},
		{[]int{1, 2, 3, 5}, false},
	}
	for _, v := range e {
		o := canPartition(v.nums)
		if o != v.expected {
			t.Error(fmt.Sprintf("%v %v", v, o))
		}
	}
}
