package Num_278_FirstBadVersion

import (
	"fmt"
	"testing"
)

func TestNum_278_FirstBadVersion(t *testing.T) {
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
