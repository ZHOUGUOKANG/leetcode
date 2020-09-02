package Num_46_permute

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	in := []int{1, 2, 3}
	o := permute(in)
	t.Error(fmt.Sprintf("%v", o))
}

func Test2(t *testing.T) {
	in := []int{0, 1}
	o := permute(in)
	t.Error(fmt.Sprintf("%v", o))
}
