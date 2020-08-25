package Num_336_palindromePairs

import (
	"reflect"
	"testing"
)

func TestPalindromePairs(t *testing.T) {
	words := []string{"abcd", "dcba", "lls", "s", "sssll"}
	expected := [][]int{{0, 1}, {1, 0}, {3, 2}, {2, 4}}
	out := palindromePairs(words)
	if !reflect.DeepEqual(expected, out) {
		t.Error(words, expected, out)
	}
}

func TestPalindromePairs1(t *testing.T) {
	words := []string{"bat", "tab", "cat"}
	expected := [][]int{{0, 1}, {1, 0}}
	out := palindromePairs(words)
	if !reflect.DeepEqual(expected, out) {
		t.Error(words, expected, out)
	}
}

func TestPalindromePairs2(t *testing.T) {
	words := []string{"a", ""}
	expected := [][]int{{0, 1}, {1, 0}}
	out := palindromePairs(words)
	if !reflect.DeepEqual(expected, out) {
		t.Error(words, expected, out)
	}
}
func TestPalindromePairs3(t *testing.T) {
	words := []string{"a", "abc", "aba", ""}
	expected := [][]int{{0, 3}, {3, 0}, {2, 3}, {3, 2}}
	out := palindromePairs(words)
	if !reflect.DeepEqual(expected, out) {
		t.Error(words, expected, out)
	}
}
