package Num_offer_04_05_isValidBST

import "testing"

func TestIsValidBST(t *testing.T) {
	root := &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}
	out := isValidBST(root)
	if out != true {
		t.Error(root, out)
	}
}

func TestIsValidBS1(t *testing.T) {
	root := &TreeNode{Val: 5, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 6}}}
	out := isValidBST(root)
	if out != false {
		t.Error(root, out)
	}
}

func TestIsValidBS2(t *testing.T) {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 1}}
	out := isValidBST(root)
	if out != false {
		t.Error(root, out)
	}
}

func TestIsValidBS3(t *testing.T) {
	root := &TreeNode{Val: 10, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 15, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 20}}}
	out := isValidBST(root)
	if out != false {
		t.Error(root, out)
	}
}
func TestIsValidBS4(t *testing.T) {
	root := &TreeNode{Val: 3, Left: &TreeNode{Val: 1, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 2}}, Right: &TreeNode{Val: 5, Left: &TreeNode{Val: 4, Left: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 6}}}
	out := isValidBST(root)
	if out != false {
		t.Error(root, out)
	}
}
