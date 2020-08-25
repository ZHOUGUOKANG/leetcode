package Num_394_DecodeString

import (
	"testing"
)

func TestDecodeString(t *testing.T) {
	input, output, expected := "", "", ""
	input = "3[a]2[bc]"
	expected = "aaabcbc"
	output = decodeString(input)
	if output != expected {
		t.Error(input, output, expected)
	}

	input = "3[a2[c]]"
	expected = "accaccacc"
	output = decodeString(input)
	if output != expected {
		t.Error(input, output, expected)
	}

	input = "2[abc]3[cd]ef"
	expected = "abcabccdcdcdef"
	output = decodeString(input)
	if output != expected {
		t.Error(input, output, expected)
	}
}
