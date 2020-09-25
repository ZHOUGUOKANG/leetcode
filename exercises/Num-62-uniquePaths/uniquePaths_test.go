package Num_62_UniquePaths

import (
	"fmt"
	"testing"
)

func TestNum_62_UniquePaths(t *testing.T) {
	e := []struct {
		m, n     int
		expected int
	}{
		{3, 2, 3},
		{7, 3, 28},
		{23, 12, 193536720},
	}

	for _, v := range e {
		o := uniquePaths(v.m, v.n)
		if o != v.expected {
			t.Error(fmt.Sprintf("%v%v", v, o))
		}
	}
}
