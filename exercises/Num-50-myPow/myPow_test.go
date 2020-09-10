package Num_50_myPow

import (
	"fmt"
	"testing"
)

func TestNum_50_myPow(t *testing.T) {
	examples := []struct {
		x        float64
		n        int
		expected float64
	}{
		{2.00000, 10, 1024.00000},
		{2.10000, 3, 9.26100},
		{2.00000, -2, 0.25000},
	}

	for k, v := range examples {
		o := myPow(v.x, v.n)
		if o != v.expected {
			t.Error(k, fmt.Sprintf("%v %v", v, o))
		}
	}
}
