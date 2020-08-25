package Num_173_BSTIterator

import "testing"

func TestBSTIterator(t *testing.T) {
	root := &TreeNode{}
	root.Val = 7
	root.Left = &TreeNode{}
	root.Left.Val = 3
	root.Right = &TreeNode{}
	root.Right.Val = 15
	root.Right.Left = &TreeNode{}
	root.Right.Left.Val = 9
	root.Right.Right = &TreeNode{}
	root.Right.Right.Val = 20

	iterator := Constructor(root)
	next := 0
	hasNext := true
	next = iterator.Next() // 返回 3
	if next != 3 {
		t.Error("expected", 3, "get", next)
	}

	next = iterator.Next() // 返回 7
	if next != 7 {
		t.Error("expected", 7, "get", next)
	}

	hasNext = iterator.HasNext() // 返回 true
	if hasNext != true {
		t.Error("expected", true, "get", hasNext)
	}

	next = iterator.Next() // 返回 9
	if next != 9 {
		t.Error("expected", 9, "get", next)
	}

	hasNext = iterator.HasNext() // 返回 true
	if hasNext != true {
		t.Error("expected", true, "get", hasNext)
	}

	next = iterator.Next() // 返回 15
	if next != 15 {
		t.Error("expected", 15, "get", next)
	}

	hasNext = iterator.HasNext() // 返回 true
	if hasNext != true {
		t.Error("expected", true, "get", hasNext)
	}

	next = iterator.Next() // 返回 20
	if next != 20 {
		t.Error("expected", 20, "get", next)
	}

	hasNext = iterator.HasNext() // 返回 false
	if hasNext != false {
		t.Error("expected", false, "get", hasNext)
	}
}
