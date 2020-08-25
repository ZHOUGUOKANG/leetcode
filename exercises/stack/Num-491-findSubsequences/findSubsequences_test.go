package Num_491_findSubsequences

import (
	"testing"
)

func TestFindSubsequences(t *testing.T) {
	tests := []*struct {
		nums     []int
		expected map[string]bool
	}{
		{[]int{4, 7, 7, 7}, map[string]bool{"4-6": true, "4-7": true, "4-6-7": true, "4-6-7-7": true, "6-7": true, "6-7-7": true, "7-7": true, "4-7-7": true}},
	}

	for _, v := range tests {
		out := findSubsequences(v.nums)
		if len(out) != len(v.expected) {
			t.Error(v, out)
		}
		for _, o := range out {
			str := ""
			for _, z := range o {
				str += "-" + string(z)
			}
			if _, ok := v.expected[str]; ok {
				t.Error(v, out, str)
			}
		}
	}
}
