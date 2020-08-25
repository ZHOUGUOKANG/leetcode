package Num_392_isSubsequence

import "testing"

func TestIsSubsequence(t *testing.T) {
	if isSubsequence("abc", "ahbgdc") != true {
		t.Error()
	}
}

func TestIsSubsequence1(t *testing.T) {
	if isSubsequence("axc", "ahbgdc") != false {
		t.Error()
	}
}
