package Num_338_countBits

import (
	"reflect"
	"testing"
)

func TestCountBits(t *testing.T) {
	output := countBits(2)
	expected := []int{0, 1, 1}
	if !reflect.DeepEqual(output, expected) {
		t.Error(expected, output)
	}
	output = countBits(5)
	expected = []int{0, 1, 1, 2, 1, 2}
	if !reflect.DeepEqual(output, expected) {
		t.Error(expected, output)
	}
}
