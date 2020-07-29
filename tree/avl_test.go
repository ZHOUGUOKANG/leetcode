package tree

import (
	"fmt"
	"testing"
)

func TestAvl(t *testing.T) {
	root := NewTree()
	root.Insert(5)
	root.Insert(6)
	root.Insert(10)
	root.Insert(7)
	root.Insert(3)
	root.Insert(15)
	root.Insert(20)
	root.Insert(33)
	root.Insert(1)
	root.Insert(8)
	root.Insert(11)
	root.Insert(12)
	root.Insert(100)
	root.Insert(102)
	root.Insert(103)
	println(fmt.Sprintf("%s", root))
}
