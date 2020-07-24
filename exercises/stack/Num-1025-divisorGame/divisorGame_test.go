package Num_1025_divisorGame

import "testing"

func TestDivisorGame(t *testing.T) {
	input := 2
	expected := true
	output := divisorGame(input)
	if output != expected {
		t.Error(input, expected, output)
	}
}

func TestDivisorGame1(t *testing.T) {
	input := 3
	expected := false
	output := divisorGame(input)
	if output != expected {
		t.Error(input, expected, output)
	}
}
