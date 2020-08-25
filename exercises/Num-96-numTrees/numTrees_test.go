package Num_96_numTrees

import "testing"

func TestNumTrees(t *testing.T) {
	expected := 1
	input := 0
	output := numTrees(input)
	if expected != output {
		t.Error(expected, output)
	}

	expected = 2
	input = 2
	output = numTrees(input)
	if expected != output {
		t.Error(expected, output)
	}

	expected = 5
	input = 3
	output = numTrees(input)
	if expected != output {
		t.Error(expected, output)
	}
}
