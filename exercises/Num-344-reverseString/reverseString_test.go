package Num_344_ReverseString

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNum_344_ReverseString(t *testing.T) {
	e := []struct {
		str      []byte
		expected []byte
	}{
		{[]byte{}, []byte{}},
		{[]byte{'a', 'b', 'c'}, []byte{'c', 'b', 'a'}},
		{[]byte{'a', 'b', 'c', 'd'}, []byte{'d', 'c', 'b', 'a'}},
	}
	for _, v := range e {
		reverseString(v.str)
		if !reflect.DeepEqual(v.expected, v.str) {
			t.Error(fmt.Sprintf("%v", v))
		}
	}
}
