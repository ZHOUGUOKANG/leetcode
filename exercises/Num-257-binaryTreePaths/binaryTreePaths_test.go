package Num_257_binaryTreePaths

import (
	"fmt"
	"math"
	"testing"
)

func TestNum_257_binaryTreePaths(t *testing.T) {
	examples := []struct {
		nodes    []int
		expected []string
	}{
		{[]int{1, 2, 3, math.MinInt32, 5}, []string{"1->2->5", "1->3"}},
	}

	for _, v := range examples {
		nodes := make([]*TreeNode, len(v.nodes))
		for i := 0; i < len(nodes); i++ {
			if v.nodes[i] == math.MinInt32 {
				continue
			}
			nodes[i] = &TreeNode{Val: v.nodes[i]}
			rootID := (i - 1) / 2
			if i%2 == 1 {
				nodes[rootID].Left = nodes[i]
			} else {
				nodes[rootID].Right = nodes[i]
			}
		}
		out := binaryTreePaths(nodes[0])
		mmap := make(map[string]bool)
		for _, d := range out {
			mmap[d] = true
		}
		//println(fmt.Sprintf("%v %v",v,out))
		//for i:=0;i<len(nodes);i++{
		//	if nodes[i] != nil{
		//		//println(nodes[i].Val)
		//		if nodes[i].Left != nil{
		//			println("Left",nodes[i].Left.Val)
		//		}
		//		if nodes[i].Right != nil{
		//			println("Right",nodes[i].Right.Val)
		//		}
		//	}else {
		//		println("nil")
		//	}
		//}
		if len(mmap) != len(v.expected) {
			t.Error(fmt.Sprintf("%v %v", v, out))
		}
		for _, d := range out {
			if _, ok := mmap[d]; !ok {
				t.Error(fmt.Sprintf("%v %v", v, out), d)
			}
		}
	}
}
