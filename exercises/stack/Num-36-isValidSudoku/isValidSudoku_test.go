package Num_36_isValidSudoku

import "testing"

func TestIsValidSudoku(t *testing.T) {
	tests := []*struct {
		input    [][]byte
		expected bool
	}{
		{[][]byte{
			{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
			{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
			{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
			{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
			{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
			{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
			{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
			{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
			{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
		}, true},
		//{[][]byte{
		//	{'8','3','.','.','7','.','.','.','.'},
		//	{'6','.','.','1','9','5','.','.','.'},
		//	{'.','9','8','.','.','.','.','6','.'},
		//	{'8','.','.','.','6','.','.','.','3'},
		//	{'4','.','.','8','.','3','.','.','1'},
		//	{'7','.','.','.','2','.','.','.','6'},
		//	{'.','6','.','.','.','.','2','8','.'},
		//	{'.','.','.','4','1','9','.','.','5'},
		//	{'.','.','.','.','8','.','.','7','9'},
		//},false},
	}

	for _, v := range tests {
		out := isValidSudoku(v.input)
		if out != v.expected {
			t.Error(v, out)
		}
	}
}
