package Num_529_updateBoard

import (
	"reflect"
	"testing"
)

func TestUpdateBoard(t *testing.T) {
	tests := []*struct {
		input    [][]byte
		click    []int
		expected [][]byte
	}{
		{
			[][]byte{{'B', '1', 'E', '1', 'B'}, {'B', '1', 'M', '1', 'B'}, {'B', '1', '1', '1', 'B'}, {'B', 'B', 'B', 'B', 'B'}},
			[]int{1, 2},
			[][]byte{{'B', '1', 'E', '1', 'B'}, {'B', '1', 'X', '1', 'B'}, {'B', '1', '1', '1', 'B'}, {'B', 'B', 'B', 'B', 'B'}},
		},
		{
			[][]byte{{'E', 'E', 'E', 'E', 'E'}, {'E', 'E', 'M', 'E', 'E'}, {'E', 'E', 'E', 'E', 'E'}, {'E', 'E', 'E', 'E', 'E'}},
			[]int{3, 0},
			[][]byte{{'B', '1', 'E', '1', 'B'}, {'B', '1', 'M', '1', 'B'}, {'B', '1', '1', '1', 'B'}, {'B', 'B', 'B', 'B', 'B'}},
		},
	}

	for _, v := range tests {
		for i := 0; i < len(v.expected); i++ {
			println(string(v.input[i]))
		}
		out := updateBoard(v.input, v.click)
		if !reflect.DeepEqual(out, v.expected) {
			t.Error(v.click)
			for i := 0; i < len(v.expected); i++ {
				println("  ", string(v.expected[i]), "  ", string(out[i]))
			}
		}
	}
}
