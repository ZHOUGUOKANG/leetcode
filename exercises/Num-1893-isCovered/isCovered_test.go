package Num_1893_IsCovered

import (
	"fmt"
	"testing"
)

func TestNum_1893_IsCovered(t *testing.T) {

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
