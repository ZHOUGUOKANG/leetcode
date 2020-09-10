package Num_49_groupAnagrams

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNum_49_groupAnagrams(t *testing.T) {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	expected := [][]string{
		{"ate", "eat", "tea"},
		{"nat", "tan"},
		{"bat"},
	}
	o := groupAnagrams(strs)
	if !reflect.DeepEqual(expected, o) {
		t.Error(fmt.Sprintf("%v%v%v", strs, expected, o))
	}
}
