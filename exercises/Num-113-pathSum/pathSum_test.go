package Num_113_PathSum

import "testing"

func TestNum_113_PathSum(t *testing.T) {

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
