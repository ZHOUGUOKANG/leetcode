package Num_297_Serialize

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

type serializeData [][]int

// Serializes a tree to a single string.
func (c *Codec) serialize(root *TreeNode) string {
	b := strings.Builder{}
	var helper func(n *TreeNode)
	helper = func(n *TreeNode) {
		if n == nil {
			b.WriteString("nil,")
			return
		}
		b.WriteString(strconv.Itoa(n.Val))
		b.WriteByte(',')
		helper(n.Left)
		helper(n.Right)
	}
	helper(root)
	return b.String()
}

// Deserializes your encoded data to tree.
func (c *Codec) deserialize(data string) *TreeNode {
	arr := strings.Split(data, ",")
	var helper func() *TreeNode
	helper = func() *TreeNode {
		if len(arr) == 0 {
			return nil
		}
		if arr[0] == "nil" {
			arr = arr[1:]
			return nil
		}
		i, _ := strconv.Atoi(arr[0])
		arr = arr[1:]
		return &TreeNode{Val: i, Left: helper(), Right: helper()}
	}
	return helper()
}
