package Num_350_intersect

import (
	"reflect"
	"testing"
)

func TestIntersect(t *testing.T) {
	expected := []int{2, 2}
	output := intersect([]int{1, 2, 2, 1}, []int{2, 2})
	if !reflect.DeepEqual(expected, output) {
		t.Error(expected, output)
	}

	expected = []int{4, 9}
	output = intersect([]int{4, 9, 5}, []int{9, 4, 9, 8, 4})
	if !reflect.DeepEqual(expected, output) {
		t.Error(expected, output)
	}
}
