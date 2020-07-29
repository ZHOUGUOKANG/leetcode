package Num_378_kthSmallest

import "testing"

func TestKthSmallest(t *testing.T) {
	in, k := [][]int{{1, 5, 9}, {10, 11, 13}, {12, 13, 15}}, 8
	expected := 13
	out := kthSmallest(in, k)
	if out != expected {
		t.Error(in, k, expected, out)
	}
}
