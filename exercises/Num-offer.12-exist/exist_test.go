package Num_offer_12_exist

import (
	"fmt"
	"testing"
)

func TestExist(t *testing.T) {
	examples := []struct {
		board    [][]byte
		word     string
		expected bool
	}{
		{[][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCCED", true},
		{[][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABFS", true},
		{[][]byte{{'a', 'b'}, {'c', 'd'}}, "abcd", false},
		{[][]byte{{'a', 'b'}, {'c', 'd'}}, "bacd", true},
	}
	for k, v := range examples {
		if exist(v.board, v.word) != v.expected {
			t.Error(k, fmt.Sprintf("%v", v))
		}
	}
}
