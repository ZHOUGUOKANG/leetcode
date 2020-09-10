package Num_60_getPermutation

import (
	"fmt"
	"testing"
)

func TestNum_60_getPermutation(t *testing.T) {
	examples := []struct {
		n, k     int
		expected string
	}{
		{3, 3, "213"},
		{4, 9, "2314"},
		{9, 2678, "126847395"},
	}

	for k, v := range examples {
		o := getPermutation(v.n, v.k)
		if o != v.expected {
			t.Error(k, fmt.Sprintf("%v", v), o)
		}
	}
}
