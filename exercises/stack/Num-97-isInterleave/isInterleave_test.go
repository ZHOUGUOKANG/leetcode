package Num_97_isInterleave

import "testing"

func TestIsInterleave(t *testing.T) {
	input := []string{"aabcc", "dbbca", "aadbbcbcac"}
	expected := true
	output := isInterleave(input[0], input[1], input[2])
	if expected != output {
		t.Error(input, expected, output)
	}
}
func TestIsInterleave1(t *testing.T) {
	input := []string{"aabcc", "dbbca", "aadbbbaccc"}
	expected := false
	output := isInterleave(input[0], input[1], input[2])
	if expected != output {
		t.Error(input, expected, output)
	}
}
