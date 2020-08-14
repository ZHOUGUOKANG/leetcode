package Num_43_multiply

import (
	"reflect"
	"testing"
)

type result struct {
	num1     string
	num2     string
	expected string
}

func TestBytesAdd(t *testing.T) {
	a, b := []byte{'1'}, []byte{'9'}
	expected := []byte{'1', '0'}
	out := bytesAdd(a, b)
	if !reflect.DeepEqual(expected, out) {
		t.Error(a, b, expected, out)
	}
}
func TestBytesAdd1(t *testing.T) {
	a, b := []byte{'1', '2'}, []byte{'9', '1', '1'}
	expected := []byte{'9', '2', '3'}
	out := bytesAdd(a, b)
	if !reflect.DeepEqual(expected, out) {
		t.Error(a, b, expected, out)
	}
}
func TestBytesAdd2(t *testing.T) {
	a, b := []byte{'9', '1'}, []byte{'9'}
	expected := []byte{'1', '0', '0'}
	out := bytesAdd(a, b)
	if !reflect.DeepEqual(expected, out) {
		t.Error(a, b, expected, out)
	}
}

func TestHelper(t *testing.T) {
	a, b := uint8(9), []byte{'9'}
	expected := []byte{'8', '1'}
	out := helper(a, b)
	if !reflect.DeepEqual(expected, out) {
		t.Error(a, b, expected, out)
	}
}
func TestHelper1(t *testing.T) {
	a, b := uint8(9), []byte{'9', '1', '0'}
	expected := []byte{'8', '1', '9', '0'}
	out := helper(a, b)
	if !reflect.DeepEqual(expected, out) {
		t.Error(a, b, expected, out)
	}
}
func TestHelper2(t *testing.T) {
	a, b := uint8(0), []byte{'9'}
	out := helper(a, b)
	if out != nil {
		t.Error(a, b, nil, out)
	}
}
func TestMultiply(t *testing.T) {
	results := []*result{
		{"2", "3", "6"},
		{"0", "3", "0"},
		{"110", "90", "9900"},
		{"90", "110", "9900"},
		{"123", "456", "56088"},
		{"9124", "21250973", "193893877652"},
	}
	for _, v := range results {
		r := multiply(v.num1, v.num2)
		if r != v.expected {
			t.Error(v, r)
		}
	}
}
