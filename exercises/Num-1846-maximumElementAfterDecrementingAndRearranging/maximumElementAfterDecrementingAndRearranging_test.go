package Num_1846_MaximumElementAfterDecrementingAndRearranging

import (
	"fmt"
	"testing"
)

func TestNum_1846_MaximumElementAfterDecrementingAndRearranging(t *testing.T) {

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
