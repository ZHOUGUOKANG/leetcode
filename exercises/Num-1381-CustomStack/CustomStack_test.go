package Num_1381_CustomStack

import (
	"reflect"
	"testing"
)

func TestCustomStack(t *testing.T) {
	cs := Constructor(3)
	if cs.Pop() != -1 {
		t.Error("empty CustomStack pop must -1")
	}
	cs.Push(10)
	expected := []int{10}
	if !reflect.DeepEqual(cs.GetList(), expected) {
		t.Error("expected", expected, "get", cs.GetList())
	}

	cs.Push(20)
	expected = []int{10, 20}
	if !reflect.DeepEqual(cs.GetList(), expected) {
		t.Error("expected", expected, "get", cs.GetList())
	}

	cs.Push(30)
	expected = []int{10, 20, 30}
	if !reflect.DeepEqual(cs.GetList(), expected) {
		t.Error("expected", expected, "get", cs.GetList())
	}

	cs.Push(40)
	expected = []int{10, 20, 30}
	if !reflect.DeepEqual(cs.GetList(), expected) {
		t.Error("expected", expected, "get", cs.GetList())
	}

	value := cs.Pop()
	expected = []int{10, 20}
	if !reflect.DeepEqual(cs.GetList(), expected) {
		t.Error("expected", expected, "get", cs.GetList())
	}
	if !reflect.DeepEqual(value, 30) {
		t.Error("expected", 30, "get", value)
	}

	value = cs.Pop()
	expected = []int{10}
	if !reflect.DeepEqual(cs.GetList(), expected) {
		t.Error("expected", expected, "get", cs.GetList())
	}
	if !reflect.DeepEqual(value, 20) {
		t.Error("expected", 20, "get", value)
	}

	value = cs.Pop()
	expected = []int{}
	if !reflect.DeepEqual(cs.GetList(), expected) {
		t.Error("expected", expected, "get", cs.GetList())
	}
	if !reflect.DeepEqual(value, 10) {
		t.Error("expected", 10, "get", value)
	}

	value = cs.Pop()
	expected = []int{}
	if !reflect.DeepEqual(cs.GetList(), expected) {
		t.Error("expected", expected, "get", cs.GetList())
	}
	if !reflect.DeepEqual(value, -1) {
		t.Error("expected", -1, "get", value)
	}
}
