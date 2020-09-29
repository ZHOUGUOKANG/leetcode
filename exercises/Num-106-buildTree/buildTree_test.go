package Num_106_BuildTree

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNum_106_BuildTree(t *testing.T) {
	e := []struct {
		nums1    []int
		nums2    []int
		expected []int
	}{
		{[]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3}, []int{3, 9, 20, 15, 7}},
	}
	for _, v := range e {
		o := buildTree(v.nums1, v.nums2)
		ans := make([]int, 0)
		helper := func(n *TreeNode) {}
		helper = func(n *TreeNode) {
			if n == nil {
				return
			}
			ans = append(ans, n.Val)
			helper(n.Left)
			helper(n.Right)
		}
		helper(o)
		if !reflect.DeepEqual(v.expected, ans) {
			t.Error(fmt.Sprintf("%v %v", v, ans))
		}
	}
}
