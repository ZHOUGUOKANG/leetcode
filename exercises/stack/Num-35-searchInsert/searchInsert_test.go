package Num_35_searchInsert

import "testing"

func TestSearchInsert(t *testing.T) {
	expected := 2
	out := searchInsert([]int{1, 3, 5, 6}, 5)
	if out != expected {
		t.Error(expected, out)
	}
}

func TestSearchInsert1(t *testing.T) {
	expected := 1
	out := searchInsert([]int{1, 3, 5, 6}, 2)
	if out != expected {
		t.Error(expected, out)
	}
}
func TestSearchInsert2(t *testing.T) {
	expected := 4
	out := searchInsert([]int{1, 3, 5, 6}, 7)
	if out != expected {
		t.Error(expected, out)
	}
}
func TestSearchInsert3(t *testing.T) {
	expected := 0
	out := searchInsert([]int{1, 3, 5, 6}, 0)
	if out != expected {
		t.Error(expected, out)
	}
}
func TestSearchInsert4(t *testing.T) {
	expected := 0
	out := searchInsert([]int{1}, 0)
	if out != expected {
		t.Error(expected, out)
	}
}
func TestSearchInsert5(t *testing.T) {
	expected := 3
	out := searchInsert([]int{3, 5, 7, 9, 10}, 8)
	if out != expected {
		t.Error(expected, out)
	}
}

func TestSearchInsert6(t *testing.T) {
	expected := 1
	out := searchInsert([]int{1, 3}, 2)
	if out != expected {
		t.Error(expected, out)
	}
}

func TestSearchInsert7(t *testing.T) {
	expected := 0
	out := searchInsert([]int{1, 3, 5}, 1)
	if out != expected {
		t.Error(expected, out)
	}
}
