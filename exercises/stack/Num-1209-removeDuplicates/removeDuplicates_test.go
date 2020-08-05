package Num_1209_removeDuplicates

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	s, k := "abcd", 2
	expected := "abcd"
	out := removeDuplicates(s, k)
	if expected != out {
		t.Error(s, k, expected, out)
	}
}

func TestRemoveDuplicates1(t *testing.T) {
	s, k := "deeedbbcccbdaa", 3
	expected := "aa"
	out := removeDuplicates(s, k)
	if expected != out {
		t.Error(s, k, expected, out)
	}
}

func TestRemoveDuplicates2(t *testing.T) {
	s, k := "pbbcggttciiippooaais", 2
	expected := "ps"
	out := removeDuplicates(s, k)
	if expected != out {
		t.Error(s, k, expected, out)
	}
}
func TestRemoveDuplicates3(t *testing.T) {
	s, k := "yfttttfbbbbnnnnffbgffffgbbbbgssssgthyyyy", 4
	expected := "ybth"
	out := removeDuplicates(s, k)
	if expected != out {
		t.Error(s, k, expected, out)
	}
}

func TestRemoveDuplicates4(t *testing.T) {
	s, k := "iiiixxxxxiiccccczzffffflllllllllfffffllyyyyyuuuuuz", 5
	expected := "izzlz"
	out := removeDuplicates(s, k)
	if expected != out {
		t.Error(s, k, expected, out)
	}
}
