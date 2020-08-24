package Num_30_findSubstring

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFindSubstring(t *testing.T) {
	tests := []*struct {
		s        string
		words    []string
		expected []int
	}{
		//{"barfoothefoobarman",[]string{"foo","bar"},[]int{0,9}},
		//{"wordgoodgoodgoodbestword",[]string{"word","good","best","word"},nil},
		{"wordgoodgoodgoodbestword", []string{"word", "good", "best", "good"}, []int{8}},
	}

	for _, v := range tests {
		out := findSubstring(v.s, v.words)
		if !reflect.DeepEqual(out, v.expected) {
			t.Error(fmt.Sprintf("%v %v %v %v", v.s, v.words, v.expected, out))
		}
	}
}
