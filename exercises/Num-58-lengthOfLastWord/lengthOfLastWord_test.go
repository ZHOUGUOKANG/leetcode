package Num_58_LengthOfLastWord

import "testing"

func TestNum_58_LengthOfLastWord(t *testing.T) {
	examples := []struct {
		word     string
		expected int
	}{
		{"Hello World", 5},
		{"", 0},
		{" ", 0},
		{"word ", 4},
		{"a", 1},
	}

	for _, v := range examples {
		o := lengthOfLastWord(v.word)
		if o != v.expected {
			t.Error(v.word, v.expected, o)
		}
	}
}
