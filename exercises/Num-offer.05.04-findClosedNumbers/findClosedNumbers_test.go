package Num_offer_05_04_findClosedNumbers

import (
	"reflect"
	"testing"
)

func TestFindClosedNumbers(t *testing.T) {
	num := 2
	expected := []int{4, 1}
	out := findClosedNumbers(num)
	if !reflect.DeepEqual(expected, out) {
		t.Error(num, expected, out)
	}
}

func TestFindClosedNumbers1(t *testing.T) {
	num := 1
	expected := []int{2, -1}
	out := findClosedNumbers(num)
	if !reflect.DeepEqual(expected, out) {
		t.Error(num, expected, out)
	}
}
func TestFindClosedNumbers2(t *testing.T) {
	num := 1837591841
	expected := []int{1837591842, 1837591832}
	out := findClosedNumbers(num)
	if !reflect.DeepEqual(expected, out) {
		t.Error(num, expected, out)
	}
}
