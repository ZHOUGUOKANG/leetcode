package Num_41_firstMissingPositive

import (
	"fmt"
	"testing"
)

type testFirstMissingPositive struct {
	nums     []int
	expected int
}

func (t *testFirstMissingPositive) String() string {
	return fmt.Sprintf("input:%v expected:%d", t.nums, t.expected)
}
func TestFirstMissingPositive(t *testing.T) {
	tests := []*testFirstMissingPositive{
		{[]int{1, 2, 0}, 3},
		{[]int{1, 2, 4, 0}, 3},
		{[]int{3, 4, -1, 1}, 2},
		{[]int{7, 8, 9, 10, 11, 12}, 1},
	}

	for _, v := range tests {
		o := firstMissingPositive(v.nums)
		if o != v.expected {
			t.Error(v, " output:"+string(0))
		}
	}
}
