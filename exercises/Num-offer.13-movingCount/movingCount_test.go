package Num_offer_13_movingCount

import (
	"fmt"
	"testing"
)

func TestMovingCount(t *testing.T) {
	examples := []struct {
		m        int
		n        int
		k        int
		expected int
	}{
		{2, 3, 1, 3},
		{3, 1, 0, 1},
	}

	for k, v := range examples {
		o := movingCount(v.m, v.n, v.k)
		if o != v.expected {
			t.Error(k, fmt.Sprintf("%v %v", v, o))
		}
	}
}
