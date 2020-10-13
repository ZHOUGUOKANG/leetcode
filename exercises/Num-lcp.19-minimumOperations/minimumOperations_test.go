package Num_19_MinimumOperations

import (
	"fmt"
	"testing"
)

func TestNum_19_MinimumOperations(t *testing.T) {

	e := []struct {
		st       string
		expected int
	}{
		{"rrryyyrryyyrr", 2},
		{"ryr", 0},
	}
	for _, v := range e {
		o := minimumOperations(v.st)
		if o != v.expected {
			t.Error(fmt.Sprintf("%v %v", v, o))
		}
	}
}
