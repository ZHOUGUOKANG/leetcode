package Num_117_Connect

import (
	"fmt"
	"testing"
)

func TestNum_117_Connect(t *testing.T) {
	e := []struct {
		nums []int
	}{
		{[]int{1, 2, 3, 4, 5, -1, 7}},
	}
	for _, v := range e {
		nodes := make([]*Node, len(v.nums))
		for i, val := range v.nums {
			if val != -1 {
				nodes[i] = &Node{Val: val}
			}
			parent := nodes[(i-1)/2]
			if parent != nil {
				if i%2 == 0 {
					parent.Right = nodes[i]
				} else {
					parent.Left = nodes[i]
				}
			}
		}
		o := connect(nodes[0])
		if o.Val != 1 || o.Next != nil {
			t.Error(fmt.Sprintf("%v %v", v, o))
		}
		if o.Left.Val != 2 || o.Left.Next != o.Right || o.Right.Val != 3 {
			t.Error(fmt.Sprintf("%v %v", v, o))
		}
		if o.Left.Left.Next != o.Left.Right || o.Left.Left.Val != 4 || o.Left.Right.Val != 5 {
			t.Error(fmt.Sprintf("%v %v", v, o))
		}
		if o.Left.Right.Next != o.Right.Right || o.Left.Right.Val != 5 || o.Right.Right.Val != 7 {
			t.Error(fmt.Sprintf("%v %v", v, o))
		}
	}
}
