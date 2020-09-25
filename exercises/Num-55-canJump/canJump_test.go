package Num_55_canJump

import (
	"fmt"
	"testing"
)

func TestNum_55_canJump(t *testing.T) {
	vv := []struct {
		nums []int
		e    bool
	}{
		{[]int{2, 3, 1, 1, 4}, true},
		{[]int{3, 2, 1, 0, 4}, false},
		{[]int{1, 1, 1, 0}, true},
	}

	for _, v := range vv {
		if v.e != canJump(v.nums) {
			t.Error(fmt.Sprintf("%v", v))
		}
	}
}
