package Num_65_IsNumber

import (
	"fmt"
	"testing"
)

func TestNum_65_IsNumber(t *testing.T) {
	e := []struct {
		s        string
		expected bool
	}{
		{"0", true},
		{" 0.1", true},
		{"abc", false},
		{"1 a", false},
		{"2e10", true},
		{" -90e3   ", true},
		{" le", false},
		{"e3", false},
		{" 6e-1", true},
		{"53.5e93", true},
		{" 99e2.5", false},
		{" --6 ", false},
		{"-+3", false},
		{"95a54e53", false},
	}
	for _, v := range e {
		o := isNumber(v.s)
		if o != v.expected {
			t.Error(fmt.Sprintf("%v %v", v, o))
		}
	}
}
