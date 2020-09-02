package Num_841_canVisitAllRooms

import (
	"fmt"
	"testing"
)

func TestCanVisitAllRooms(t *testing.T) {
	examples := []struct {
		rooms    [][]int
		expected bool
	}{
		{[][]int{{1}, {2}, {3}, {}}, true},
		{[][]int{{1, 3}, {3, 0, 1}, {2}, {}}, false},
	}

	for _, e := range examples {
		o := canVisitAllRooms(e.rooms)
		if o != e.expected {
			t.Error(fmt.Sprintf("%v", e), o)
		}
	}
}
