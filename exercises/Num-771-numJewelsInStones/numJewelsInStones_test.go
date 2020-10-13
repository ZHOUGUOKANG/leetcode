package Num_771_NumJewelsInStones

import (
	"fmt"
	"testing"
)

func TestNum_771_NumJewelsInStones(t *testing.T) {

	e := []struct {
		J, S     string
		expected int
	}{
		{"aA", "aAAbbbBB", 3},
	}
	for _, v := range e {
		o := numJewelsInStones(v.J, v.S)
		if o != v.expected {
			t.Error(fmt.Sprintf("%v %v", v, o))
		}
	}
}
