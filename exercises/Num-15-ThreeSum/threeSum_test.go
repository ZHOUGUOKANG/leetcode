package Num_15_ThreeSum

import (
	"fmt"
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {
	input := []int{-1, 0, 1, 2, -1, -4}
	output := threeSum(input)
	expected := [][]int{[]int{-1, 0, 1}, []int{-1, -1, 2}}
	if !reflect.DeepEqual(output, expected) {
		t.Error(input, output, expected)
	}

	input = []int{0, 0, 0}
	output = threeSum(input)
	expected = [][]int{[]int{0, 0, 0}}
	println(fmt.Sprintf("%v,%v,%v", input, output, expected))
	if !reflect.DeepEqual(output, expected) {
		t.Error(input, output, expected)
	}
}
