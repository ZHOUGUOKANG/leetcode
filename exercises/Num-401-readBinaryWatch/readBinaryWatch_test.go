package Num_401_ReadBinaryWatch

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestNum_401_ReadBinaryWatch(t *testing.T) {
	e := []struct {
		input    int
		expected []string
	}{
		{input: 1, expected: []string{"0:01", "0:02", "0:04", "0:08", "0:16", "0:32", "1:00", "2:00", "4:00", "8:00"}},
		{input: 0, expected: []string{"0:00"}},
	}
	for _, v := range e {
		o := readBinaryWatch(v.input)
		sort.Strings(o)
		sort.Strings(v.expected)
		if !reflect.DeepEqual(o, v.expected) {
			t.Error(fmt.Sprintf("%v %v", v, o))
		}
	}
}
