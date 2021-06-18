package Num_483_SmallestGoodBase

import (
	"fmt"
	"testing"
)

func TestNum_483_SmallestGoodBase(t *testing.T) {
	e := []struct {
		input    string
		expected string
	}{
		{"13", "3"},
		{"3", "2"},
		{"4681", "8"},
		{"1000000000000000000", "999999999999999999"},
		{"2251799813685247", "2"},
	}
	for _, v := range e {
		o := smallestGoodBase(v.input)
		if o != v.expected {
			t.Error(fmt.Sprintf("%v %v", v, o))
		}
	}
}
