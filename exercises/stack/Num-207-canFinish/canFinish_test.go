package Num_207_canFinish

import "testing"

func TestCanFinish(t *testing.T) {
	c, p := 2, [][]int{{1, 0}}
	expected := true
	out := canFinish(c, p)
	if out != expected {
		t.Error(c, p, expected, out)
	}
}

func TestCanFinish1(t *testing.T) {
	c, p := 2, [][]int{{1, 0}, {0, 1}}
	expected := false
	out := canFinish(c, p)
	if out != expected {
		t.Error(c, p, expected, out)
	}
}
