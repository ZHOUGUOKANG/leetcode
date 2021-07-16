package Num_981_TimeMap

import (
	"fmt"
	"testing"
)

func TestNum_981_TimeMap(t *testing.T) {

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
