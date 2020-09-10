package Num_107_levelOrderBottom

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

func TestNum_107_levelOrderBottom(t *testing.T) {
	examples := []struct {
		nodes    []int
		expected [][]int
	}{
		{[]int{3, 9, 20, math.MinInt32, math.MinInt32, 15, 7}, [][]int{{15, 7}, {9, 20}, {3}}},
	}

	for k, v := range examples {
		nodes := make([]*TreeNode, len(v.nodes))
		for index, val := range v.nodes {
			if val == math.MinInt32 {
				continue
			}
			nodes[index] = &TreeNode{Val: val}
			fNodeID := (index - 1) / 2
			if index%2 == 0 {
				nodes[fNodeID].Right = nodes[index]
			} else {
				nodes[fNodeID].Left = nodes[index]
			}
		}
		o := levelOrderBottom(nodes[0])
		if !reflect.DeepEqual(o, v.expected) {
			t.Error(k, fmt.Sprintf("%v %v", v.expected, o))
		}
	}
}
