package Num_100_isSameTree

import "testing"

func TestIsSameTree(t *testing.T) {
	r1 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	r2 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	out := isSameTree(r1, r2)
	if out != true {
		t.Error(true, out)
	}
}

func TestIsSameTree1(t *testing.T) {
	r1 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}
	r2 := &TreeNode{Val: 1, Right: &TreeNode{Val: 2}}
	out := isSameTree(r1, r2)
	if out != false {
		t.Error(false, out)
	}
}
