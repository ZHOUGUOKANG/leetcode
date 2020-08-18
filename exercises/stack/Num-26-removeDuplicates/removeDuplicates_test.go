package Num_26_removeDuplicates

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	in1 := []int{1, 1, 2}
	in2 := []int{1, 2}
	expected := 2
	out := removeDuplicates(in1)
	if out != expected || !reflect.DeepEqual(in1[:expected], in2) {
		t.Error(in2, expected, in1, out)
	}
}

func TestRemoveDuplicates1(t *testing.T) {
	in1 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	in2 := []int{0, 1, 2, 3, 4}
	expected := 5
	out := removeDuplicates(in1)
	if out != expected || !reflect.DeepEqual(in1[:expected], in2[:expected]) {
		t.Error(in2, expected, in1, out)
	}
}
