package Num_103_ZigzagLevelOrder

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNum_103_ZigzagLevelOrder(t *testing.T) {
	e := []struct {
		nums     []int
		expected [][]int
	}{
		{[]int{3, 9, 20, -1, -1, 15, 7}, [][]int{{3}, {20, 9}, {15, 7}}},
		{[]int{1, 2, 3, 4, -1, -1, 5}, [][]int{{1}, {3, 2}, {4, 5}}},
	}
	for _, v := range e {
		root := BuildTree(v.nums)
		o := zigzagLevelOrder(root)
		if !reflect.DeepEqual(v.expected, o) {
			t.Error(fmt.Sprintf("%v %v", v, o))
		}
	}
}

func BuildTree(nums []int) *TreeNode {
	nodes := make([]*TreeNode, len(nums))
	for k, v := range nums {
		if v == -1 {
			nodes[k] = nil
		} else {
			nodes[k] = &TreeNode{Val: v}
		}
		parent := nodes[(k-1)/2]
		if parent != nil && k > 0 {
			if k%2 == 0 {
				parent.Right = nodes[k]
			} else {
				parent.Left = nodes[k]
			}
		}
	}
	return nodes[0]
}

func PreorderTraversal(root *TreeNode) []int {
	ans := make([]int, 0)
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		pop := stack[0]
		stack = stack[1:]
		if pop == nil {
			ans = append(ans, -1)
			continue
		}
		ans = append(ans, pop.Val)
		stack = append(stack, pop.Left, pop.Right)
	}
	return ans
}
