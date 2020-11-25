package Num_1370_SortString

import (
	"fmt"
	"testing"
)

func TestNum_1370_SortString(t *testing.T) {
	e := []struct {
		s, e string
	}{
		{"aaaabbbbcccc", "abccbaabccba"},
		{"rat", "art"},
		{"leetcode", "cdelotee"},
		{"ggggggg", "ggggggg"},
		{"spo", "ops"},
	}
	for _, v := range e {
		o := sortString(v.s)
		if o != v.e {
			t.Error(fmt.Sprintf("%v %v", v, o))
		}
	}
}
