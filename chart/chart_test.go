package chart

import "testing"

func TestShortestPath(t *testing.T) {
	input := map[string]map[string]int{
		"A": map[string]int{"B": 2, "C": 1},
		"B": map[string]int{"A": 2, "D": 3},
		"C": map[string]int{"A": 1, "D": 2, "F": 3},
		"D": map[string]int{"B": 3, "C": 2, "E": 1},
		"E": map[string]int{"D": 1, "F": 1},
		"F": map[string]int{"C": 3, "E": 1},
	}
	path, length := ShortestPath(input, "A", "F")
	if path != "A-C-F" && length != 4 {
		t.Error(path, length, "A-C-F", 4)
	}

	path, length = ShortestPath(input, "A", "C")
	if path != "A-C-F" && length != 1 {
		t.Error(path, length, "A-C", 1)
	}

	path, length = ShortestPath(input, "A", "E")
	if path != "A-C-D-E" && length != 4 {
		t.Error(path, length, "A-C-D-E", 4)
	}
}
