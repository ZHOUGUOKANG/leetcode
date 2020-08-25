package Num_39_combinationSum

import (
	"fmt"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	tests := []*struct {
		candidates []int
		target     int
		expected   [][]int
	}{
		{[]int{2, 3, 6, 7}, 7, [][]int{{7}, {2, 2, 3}}},
		{[]int{2, 3, 5}, 8, [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}},
	}

	for _, v := range tests {
		out := combinationSum(v.candidates, v.target)
		if len(out) != len(v.expected) {
			t.Fatal(fmt.Sprintf("%v %v expected::%v out::%v", v.candidates, v.target, v.expected, out))
		}
		outMap := make(map[string]bool)
		for _, z := range out {
			str := ""
			for _, x := range z {
				str += "-" + string(x)
			}
			outMap[str] = true
		}
		for _, z := range v.expected {
			str := ""
			for _, x := range z {
				str += "-" + string(x)
			}
			if _, ok := outMap[str]; !ok {
				t.Fatal(fmt.Sprintf("%v %v expected::%v out::%v", v.candidates, v.target, v.expected, out))

			}
		}
	}
}
