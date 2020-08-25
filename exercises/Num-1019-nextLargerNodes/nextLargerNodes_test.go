package Num_1019_nextLargerNodes

import (
	"reflect"
	"testing"
)

func TestNextLargerNodes(t *testing.T) {
	var next *ListNode
	head := &ListNode{Val: 2}
	next = head

	next.Next = &ListNode{Val: 1}
	next = next.Next

	next.Next = &ListNode{Val: 5}
	next = next.Next

	out := nextLargerNodes(head)
	expected := []int{5, 5, 0}
	if !reflect.DeepEqual(out, expected) {
		t.Error(expected, out)
	}
}

func TestNextLargerNodes1(t *testing.T) {
	var next *ListNode
	head := &ListNode{Val: 2}
	next = head

	next.Next = &ListNode{Val: 7}
	next = next.Next

	next.Next = &ListNode{Val: 4}
	next = next.Next

	next.Next = &ListNode{Val: 3}
	next = next.Next

	next.Next = &ListNode{Val: 5}
	next = next.Next

	out := nextLargerNodes(head)
	expected := []int{7, 0, 5, 5, 0}
	if !reflect.DeepEqual(out, expected) {
		t.Error(expected, out)
	}
}

func TestNextLargerNodes2(t *testing.T) {
	var next *ListNode
	head := &ListNode{Val: 1}
	next = head

	next.Next = &ListNode{Val: 7}
	next = next.Next

	next.Next = &ListNode{Val: 5}
	next = next.Next

	next.Next = &ListNode{Val: 1}
	next = next.Next

	next.Next = &ListNode{Val: 9}
	next = next.Next

	next.Next = &ListNode{Val: 2}
	next = next.Next

	next.Next = &ListNode{Val: 5}
	next = next.Next

	next.Next = &ListNode{Val: 1}
	next = next.Next

	out := nextLargerNodes(head)
	expected := []int{7, 9, 9, 9, 0, 5, 0, 0}
	if !reflect.DeepEqual(out, expected) {
		t.Error(expected, out)
	}
}

func TestNextLargerNodes3(t *testing.T) {
	var next *ListNode
	head := &ListNode{Val: 9}
	next = head

	next.Next = &ListNode{Val: 7}
	next = next.Next

	next.Next = &ListNode{Val: 6}
	next = next.Next

	next.Next = &ListNode{Val: 7}
	next = next.Next

	next.Next = &ListNode{Val: 6}
	next = next.Next

	next.Next = &ListNode{Val: 9}
	next = next.Next

	out := nextLargerNodes(head)
	expected := []int{0, 9, 7, 9, 9, 0}
	if !reflect.DeepEqual(out, expected) {
		t.Error(expected, out)
	}
}
