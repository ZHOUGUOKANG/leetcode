package Num_1047_removeDuplicates

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	var in, expected, output string = "abbaca", "ca", ""
	output = RemoveDuplicates(in)
	if output != expected {
		t.Fatal("input", in, "expected", expected, "output", output)
	}

	in, expected = "ccabbac", "c"
	output = RemoveDuplicates(in)
	if output != expected {
		t.Fatal("input", in, "expected", expected, "output", output)
	}

	in, expected = "fggfffgb", "gb"
	output = RemoveDuplicates(in)
	if output != expected {
		t.Fatal("input", in, "expected", expected, "output", output)
	}

	in, expected = "aababaab", "ba"
	output = RemoveDuplicates(in)
	if output != expected {
		t.Fatal("input", in, "expected", expected, "output", output)
	}
	in, expected = "aaaaaaaa", ""
	output = RemoveDuplicates(in)
	if output != expected {
		t.Fatal("input", in, "expected", expected, "output", output)
	}
}
