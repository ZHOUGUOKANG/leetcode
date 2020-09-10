package Num_offer_20_isNumber

import (
	"fmt"
	"testing"
)

func TestIsNumber(t *testing.T) {
	example := []struct {
		s        string
		expected bool
	}{
		{"+100", true},
		{"5e2", true},
		{"-123", true},
		{"3.1415926", true},
		{"-3.1415926", true},
		{"-1E-16", true},
		{"0123", true},
		{"3.", true},
		{"1 ", true},
		{"1  ", true},
		{" 1", true},
		{".1", true},
		{"12e", false},
		{"1a3.14", false},
		{"1.2.3", false},
		{"+-5", false},
		{"12e+5.4", false},
		{".", false},
		{"e9", false},
		{" ", false},
		{"  ", false},
		{"1 1", false},
		{"1 .", false},
	}

	for k, v := range example {
		if isNumber(v.s) != v.expected {
			t.Error(k, fmt.Sprintf("%v", v))
		}
	}
}
