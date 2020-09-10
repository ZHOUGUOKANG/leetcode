package Num_offer_05_replaceSpace

import (
	"strings"
	"testing"
)

func TestReplaceSpace(t *testing.T) {
	example := []struct {
		str string
	}{
		{"i am a boy"},
		{"he is  a xxx"},
	}
	for k, v := range example {
		o := replaceSpace(v.str)
		expected := strings.ReplaceAll(v.str, " ", "%20")
		if o != expected {
			println(o)
			println(expected)
			t.Error(k, v.str, o, len(o), expected, len(expected))
		}
	}
}
