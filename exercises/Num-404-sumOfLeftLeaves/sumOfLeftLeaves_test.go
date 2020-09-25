package Num_404_SumOfLeftLeaves

import "testing"

func TestNum_404_SumOfLeftLeaves(t *testing.T) {
	e := []struct {
		expected interface{}
	}{
		{},
	}
	for _, v := range e {
		o := 0
		if o != v.expected {
			t.Error(fmt.Sprintf("%!v(MISSING) %!v(MISSING)", v, o))
		}
	}
}
