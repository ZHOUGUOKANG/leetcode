package Num_17_NumWays

import (
	"fmt"
	"testing"
)

func TestNum_17_NumWays(t *testing.T) {
	e := []struct {
		n, k     int
		relation [][]int
		expected int
	}{
		{n: 5, k: 3, relation: [][]int{{0, 2}, {2, 1}, {3, 4}, {2, 3}, {1, 4}, {2, 0}, {0, 4}}, expected: 3},
		{n: 3, k: 2, relation: [][]int{{0, 2}, {2, 1}}, expected: 0},
	}
	for _, v := range e {
		o := numWays(v.n, v.relation, v.k)
		if o != v.expected {
			t.Error(fmt.Sprintf("%v %v", v, o))
		}
	}
}
