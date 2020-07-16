package Num_785_isBipartite

import "testing"

func TestIsBipartite(t *testing.T) {
	out := isBipartite([][]int{{1, 3}, {0, 2}, {1, 3}, {0, 2}})
	if out != true {
		t.Error(out)
	}

	out = isBipartite([][]int{{1, 2, 3}, {0, 2}, {0, 1, 3}, {0, 2}})
	if false != out {
		t.Error(out)
	}

	if true != isBipartite([][]int{{1}, {0, 3}, {3}, {1, 2}}) {
		t.Error()
	}

	if true != isBipartite([][]int{{3}, {2, 4}, {1}, {0, 4}, {1, 3}}) {
		t.Error()
	}

	if true != isBipartite([][]int{{1}, {0}, {4}, {4}, {2}, {3}}) {
		t.Error()
	}
}
