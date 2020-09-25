package Num_67_AddBinary

import (
	"fmt"
	"testing"
)

func TestNum_67_AddBinary(t *testing.T) {
	e := []struct {
		a, b     string
		expected string
	}{
		{"11", "1", "100"},
		{"1010", "1011", "10101"},
	}

	for _, v := range e {
		o := addBinary(v.a, v.b)
		if o != v.expected {
			t.Error(fmt.Sprintf("%v", v), o)
		}
	}
}
