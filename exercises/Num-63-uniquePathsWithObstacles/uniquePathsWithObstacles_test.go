package Num_63_uniquePathsWithObstacles

import "testing"

func TestUniquePathsWithObstacles(t *testing.T) {
	input := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
	expected := 2
	output := uniquePathsWithObstacles(input)
	if expected != output {
		t.Error(input, expected, output)
	}
}

func TestUniquePathsWithObstacles1(t *testing.T) {
	input := [][]int{{1}}
	expected := 0
	output := uniquePathsWithObstacles(input)
	if expected != output {
		t.Error(input, expected, output)
	}
}

func TestUniquePathsWithObstacles2(t *testing.T) {
	input := [][]int{{0}, {1}}
	expected := 0
	output := uniquePathsWithObstacles(input)
	if expected != output {
		t.Error(input, expected, output)
	}
}
