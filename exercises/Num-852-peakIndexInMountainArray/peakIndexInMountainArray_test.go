package Num_852_PeakIndexInMountainArray

import (
	"fmt"
	"testing"
)

func TestNum_852_PeakIndexInMountainArray(t *testing.T) {
	e := []struct {
		input    []int
		expected int
	}{
		{input: []int{0, 1, 0}, expected: 1},
		{input: []int{0, 2, 1, 0}, expected: 1},
		{input: []int{0, 10, 5, 2}, expected: 1},
		{input: []int{3, 4, 5, 1}, expected: 2},
		{input: []int{24, 69, 100, 99, 79, 78, 67, 36, 26, 19}, expected: 2},
	}
	for _, v := range e {
		o := peakIndexInMountainArray(v.input)
		if o != v.expected {
			t.Error(fmt.Sprintf("%v %v", v, o))
		}
	}
}

func TestNum_852_PeakIndexInMountainArray1(t *testing.T) {
	e := []struct {
		input    []int
		expected int
	}{
		{input: []int{0, 1, 0}, expected: 1},
		{input: []int{0, 2, 1, 0}, expected: 1},
		{input: []int{0, 10, 5, 2}, expected: 1},
		{input: []int{3, 4, 5, 1}, expected: 2},
		{input: []int{24, 69, 100, 99, 79, 78, 67, 36, 26, 19}, expected: 2},
	}
	for _, v := range e {
		o := peakIndexInMountainArray1(v.input)
		if o != v.expected {
			t.Error(fmt.Sprintf("%v %v", v, o))
		}
	}
}

func TestNum_852_PeakIndexInMountainArray2(t *testing.T) {
	e := []struct {
		input    []int
		expected int
	}{
		{input: []int{0, 1, 0}, expected: 1},
		{input: []int{0, 2, 1, 0}, expected: 1},
		{input: []int{0, 10, 5, 2}, expected: 1},
		{input: []int{3, 4, 5, 1}, expected: 2},
		{input: []int{24, 69, 100, 99, 79, 78, 67, 36, 26, 19}, expected: 2},
	}
	for _, v := range e {
		o := peakIndexInMountainArray2(v.input)
		if o != v.expected {
			t.Error(fmt.Sprintf("%v %v", v, o))
		}
	}
}
