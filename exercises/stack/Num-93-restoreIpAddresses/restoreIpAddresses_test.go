package Num_93_restoreIpAddresses

import (
	"reflect"
	"testing"
)

func TestRestoreIpAddresses(t *testing.T) {
	in := "25525511135"
	expected := []string{"255.255.11.135", "255.255.111.35"}
	out := restoreIpAddresses(in)
	if !reflect.DeepEqual(expected, out) {
		t.Error(in, expected, out)
	}
}
