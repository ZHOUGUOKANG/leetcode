package Num_337_rob

import "testing"

func TestRob(t *testing.T) {
	root := &TreeNode{Val: 3, Left: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 1}}}
	expected := 7
	out := rob(root)
	if expected != out {
		t.Error(root, expected, out)
	}
}

func TestRob1(t *testing.T) {
	root := &TreeNode{Val: 3, Left: &TreeNode{Val: 4, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 5, Right: &TreeNode{Val: 1}}}
	expected := 9
	out := rob(root)
	if expected != out {
		t.Error(root, expected, out)
	}
}

func TestRob2(t *testing.T) {
	root := &TreeNode{Val: 4, Left: &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}}
	expected := 7
	out := rob(root)
	if expected != out {
		t.Error(root, expected, out)
	}
}
