package Num_332_findItinerary

import (
	"reflect"
	"testing"
)

func TestFindItinerary(t *testing.T) {
	tests := []*struct {
		tickets  [][]string
		expected []string
	}{
		{[][]string{{"MUC", "LHR"}, {"JFK", "MUC"}, {"SFO", "SJC"}, {"LHR", "SFO"}}, []string{"JFK", "MUC", "LHR", "SFO", "SJC"}},
		{[][]string{{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "ATL"}, {"ATL", "JFK"}, {"ATL", "SFO"}}, []string{"JFK", "ATL", "JFK", "SFO", "ATL", "SFO"}},
		{[][]string{{"JFK", "KUL"}, {"JFK", "NRT"}, {"NRT", "JFK"}}, []string{"JFK", "NRT", "JFK", "KUL"}},
	}

	for _, v := range tests {
		out := findItinerary(v.tickets)
		if !reflect.DeepEqual(v.expected, out) {
			t.Error(v, out)
		}
	}
}
