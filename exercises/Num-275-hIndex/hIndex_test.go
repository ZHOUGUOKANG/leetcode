package Num_275_HIndex

import (
	"fmt"
	"testing"
)

func TestNum_275_HIndex(t *testing.T) {

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
