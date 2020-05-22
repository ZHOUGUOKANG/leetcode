package sort

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	var input = []int{5, 3}
	var expected = []int{3, 5}
	output := QuickSort(input)
	if !reflect.DeepEqual(output, expected) {
		t.Error(output, expected)
	}

	input = []int{3, 5, 1}
	expected = []int{1, 3, 5}
	output = QuickSort(input)
	if !reflect.DeepEqual(output, expected) {
		t.Error(output, expected)
	}

	input = []int{3, 5, 1, 7, 17, 4, 6, 10, 7}
	expected = []int{1, 3, 4, 5, 6, 7, 7, 10, 17}
	output = QuickSort(input)
	if !reflect.DeepEqual(output, expected) {
		t.Error(output, expected)
	}
}
