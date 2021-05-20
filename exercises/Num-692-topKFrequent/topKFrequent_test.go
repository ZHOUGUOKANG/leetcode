package Num_692_TopKFrequent

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNum_692_TopKFrequent(t *testing.T) {

	e := []struct {
		words    []string
		k        int
		expected []string
	}{
		{words: []string{"i", "love", "leetcode", "i", "love", "coding"}, k: 2, expected: []string{"i", "love"}},
		{words: []string{"i", "love", "leetcode", "i", "love", "coding"}, k: 3, expected: []string{"i", "love", "coding"}},
		{words: []string{"a", "aa", "aaa"}, k: 1, expected: []string{"a"}},
	}
	for _, v := range e {
		o := topKFrequent(v.words, v.k)
		if !reflect.DeepEqual(v.expected, o) {
			t.Error(fmt.Sprintf("%v %v\n", v, o))
		}
	}
}
