package Num_752_OpenLock

import (
	"fmt"
	"testing"
)

func TestNum_752_OpenLock(t *testing.T) {
	e := []struct {
		deadNums []string
		target   string
		expected int
	}{
		{deadNums: []string{"0201", "0101", "0102", "1212", "2002"}, target: "0202", expected: 6},
		{deadNums: []string{"8888"}, target: "0009", expected: 1},
	}
	for _, v := range e {
		o := openLock(v.deadNums, v.target)
		if o != v.expected {
			t.Error(fmt.Sprintf("%v %v", v, o))
		}
	}
}
