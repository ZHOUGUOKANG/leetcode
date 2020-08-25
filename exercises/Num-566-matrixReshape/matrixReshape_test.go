package Num_566_matrixReshape

import (
	"reflect"
	"testing"
)

func TestMatrixReshape(t *testing.T) {
	nums, r, c := [][]int{{1, 2}, {3, 4}}, 1, 4
	expected := [][]int{{1, 2, 3, 4}}
	out := matrixReshape(nums, r, c)
	if !reflect.DeepEqual(expected, out) {
		t.Error(nums, r, c, expected, out)
	}
}
func TestMatrixReshape1(t *testing.T) {
	nums, r, c := [][]int{{1, 2}, {3, 4}}, 2, 4
	expected := [][]int{{1, 2}, {3, 4}}
	out := matrixReshape(nums, r, c)
	if !reflect.DeepEqual(expected, out) {
		t.Error(nums, r, c, expected, out)
	}
}
