package Num_105_BuildTree

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNum_105_BuildTree(t *testing.T) {
	e := []struct {
		nums1, nums2 []int
		expected     []int
	}{
		{[]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}, []int{3, 9, 20, 15, 7}},
	}
	for _, v := range e {
		o := buildTree(v.nums1, v.nums2)
		nodes := make([]int, 0)
		helper := func(n *TreeNode) {}
		helper = func(n *TreeNode) {
			if n == nil {
				return
			}
			nodes = append(nodes, n.Val)
			helper(n.Left)
			helper(n.Right)
		}
		helper(o)
		if !reflect.DeepEqual(v.expected, nodes[:len(v.expected)]) {
			t.Error(fmt.Sprintf("%v %v", v, nodes))
		}
	}
}
