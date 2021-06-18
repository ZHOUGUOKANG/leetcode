package Num_483_SmallestGoodBase

import (
	"fmt"
	"math"
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

func TestNum_483_SmallestGoodBase1(t *testing.T) {
	e := []struct {
		input    int
		expected int
	}{
		{13, 3},
		{3, 2},
		{4681, 8},
		{1000000000000000000, 999999999999999999},
		{2251799813685247, 2},
	}
	for _, v := range e {
		for i := 1; i < 10; i++ {
			o := math.Pow(float64(v.input), 1/float64(i))
			t.Error(fmt.Sprintf("%v %v,%v", v, i, o))
		}

	}
}
