package chart

import "testing"

func TestDijkstra(t *testing.T) {
	input := map[string]map[string]int{
		"A": {"B": 12, "G": 14, "F": 16},
		"B": {"A": 12, "F": 7, "C": 10},
		"C": {"B": 10, "F": 6, "E": 5, "D": 3},
		"D": {"C": 3, "E": 4},
		"E": {"G": 8, "F": 2, "C": 5, "D": 4},
		"F": {"A": 16, "B": 7, "C": 6, "E": 2, "G": 9},
		"G": {"A": 14, "F": 9, "E": 8},
	}
	path, length := Dijkstra(input, "A", "D")
	if path != "A-B-F-E-D" && length != 22 {
		t.Error(path, length)
	}

	path, length = Dijkstra(input, "A", "B")
	if path != "A-B" && length != 12 {
		t.Error(path, length)
	}

	path, length = Dijkstra(input, "A", "E")
	if path != "A-B-F-E" && length != 21 {
		t.Error(path, length)
	}

	path, length = Dijkstra(input, "B", "E")
	if path != "B-E" && length != 9 {
		t.Error(path, length)
	}

	path, length = Dijkstra(input, "B", "D")
	if path != "B-F-E-D" && length != 13 {
		t.Error(path, length)
	}

	path, length = Dijkstra(input, "G", "C")
	if path != "G-E-C" && length != 13 {
		t.Error(path, length)
	}

	path, length = Dijkstra(input, "F", "D")
	if path != "F-E-D" && length != 6 {
		t.Error(path, length)
	}
}
