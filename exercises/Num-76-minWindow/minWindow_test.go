package Num_76_MinWindow

import (
	"fmt"
	"testing"
)

func TestNum_76_MinWindow(t *testing.T) {
	e := []struct {
		s, t, expected string
	}{
		{"ADOBECODEBANC", "ABC", "BANC"},
		{"a", "a", "a"},
		{"a", "aa", ""},
		{"ab", "a", "a"},
	}
	for _, v := range e {
		o := minWindow(v.s, v.t)
		if o != v.expected {
			t.Error(fmt.Sprintf("%v %v", v, o))
		}
	}
}
