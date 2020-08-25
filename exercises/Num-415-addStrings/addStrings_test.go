package Num_415_addStrings

import "testing"

func TestAddStrings(t *testing.T) {
	num1, num2 := "111", "222"
	expected := "333"
	out := addStrings(num1, num2)
	if out != expected {
		t.Error(num1, num2, expected, out)
	}
}

func TestAddStrings1(t *testing.T) {
	num1, num2 := "776", "2225"
	expected := "3001"
	out := addStrings(num1, num2)
	if out != expected {
		t.Error(num1, num2, expected, out)
	}
}

func TestAddStrings2(t *testing.T) {
	num1, num2 := "1", "9"
	expected := "10"
	out := addStrings(num1, num2)
	if out != expected {
		t.Error(num1, num2, expected, out)
	}
}
