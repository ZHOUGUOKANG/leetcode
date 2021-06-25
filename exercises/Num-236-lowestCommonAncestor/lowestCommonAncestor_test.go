package Num_236_LowestCommonAncestor

import (
	"fmt"
	"testing"
)

func TestNum_236_LowestCommonAncestor(t *testing.T) {

	e := []struct {
		expected interface{}
	}{
		{},
	}
	for _, v := range e {
		o := 0
		if o != v.expected {
			t.Error(fmt.Sprintf("%v %v", v, o))
		}
	}
}
