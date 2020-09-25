package Num_617_MergeTrees

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNum_617_MergeTrees(t *testing.T) {
	e := []struct {
		tree1    []int
		tree2    []int
		expected []int
	}{
		{[]int{1, 3, 2, 5}, []int{2, 1, 3, -1, 4, -1, 7}, []int{3, 4, 5, 5, 4, -1, 7}},
	}
	for _, v := range e {
		tree1 := make([]*TreeNode, len(v.tree1))
		for k, val := range v.tree1 {
			if val != -1 {
				tree1[k] = &TreeNode{val, nil, nil}
			} else {
				tree1[k] = nil
			}
			parent := (k - 1) / 2
			if k > 0 && tree1[parent] != nil {
				if k%2 == 0 {
					tree1[parent].Right = tree1[k]
				} else {
					tree1[parent].Left = tree1[k]
				}
			}
		}
		tree2 := make([]*TreeNode, len(v.tree2))
		for k, val := range v.tree2 {
			if val != -1 {
				tree2[k] = &TreeNode{val, nil, nil}
			} else {
				tree2[k] = nil
			}
			parent := (k - 1) / 2
			if k > 0 && tree2[parent] != nil {
				if k%2 == 0 {
					tree2[parent].Right = tree2[k]
				} else {
					tree2[parent].Left = tree2[k]
				}
			}
		}
		o := mergeTrees(tree1[0], tree2[0])
		var ret []int
		queue := []*TreeNode{o}
		for len(queue) > 0 {
			head := queue[0]
			queue = queue[1:]
			if head == nil {
				ret = append(ret, -1)
				continue
			}
			ret = append(ret, head.Val)
			queue = append(queue, head.Left, head.Right)
		}
		if len(ret) < len(v.expected) || !reflect.DeepEqual(v.expected, ret[:len(v.expected)]) {
			t.Error(fmt.Sprintf("%v %v", v, ret))
		}
	}
}
