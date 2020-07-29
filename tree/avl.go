package tree

import (
	"fmt"
)

type Node struct {
	Left   *Node
	Right  *Node
	Value  int
	Height int
}

type AvlTree struct {
	root *Node
}

func NewTree() *AvlTree {
	return &AvlTree{}
}
func (n *Node) inOrderTraversal() {
	if n.Left != nil {
		n.Left.inOrderTraversal()
	}
	println(n.Value)
	if n.Right != nil {
		n.Right.inOrderTraversal()
	}
}

func (at *AvlTree) Insert(value int) *AvlTree {
	defer println("==============\n===============\n")
	if at.root == nil {
		at.root = &Node{Value: value}
		return at
	}
	node := &Node{Value: value}
	current := at.root
	var result int
	paths := make([]*Node, 0)
	flag := true
	for {
		paths = append(paths, current)
		result = node.compare(current)
		println(fmt.Sprintf("\n%d - %d = %d", value, current.Value, result))
		if result > 0 {
			println("比当前大")
			if current.Right == nil {
				println("右节点为空，插入右节点")
				current.Right = node
				break
			} else {
				println("右节点不为空，继续探索右节点")
				current = current.Right
			}
		} else if result < 0 {
			println("比当前小")
			if current.Left == nil {
				println("左节点为空，插入左节点")
				current.Left = node
				break
			} else {
				println("左节点不为空，继续探索左节点")
				current = current.Left
			}
		} else {
			println("查找到重复value")
			flag = false
			break
		}
	}
	if flag {
		println("\n------开始平衡检查--------")
		for i := len(paths) - 1; i >= 0; i-- {
			path := paths[i]
			println("当前节点为", path.Value)
			println("计算高度为", path.reHeight())
			bf := path.getBalancedFactor()
			println("平衡因子为", bf)
			if bf < -1 || bf > 1 {
				println("失去平衡")
				tmp := path.rebalanced(value)
				println("重平衡后节点为", tmp.Value)
				if i == 0 {
					println("替换为root节点")
					at.root = tmp
				} else {
					parent := paths[i-1]
					if parent.Value < value {
						println("插入父节点的右节点,父节点为", parent.Value)
						parent.Right = tmp
					} else {
						println("插入父节点的左节点,父节点为", parent.Value)
						parent.Left = tmp
					}
				}
			} else {
				println("未失去平衡")
			}
			println("---------")
		}
	}
	println("当前树结构为", at.String())
	return at
}

func (n *Node) rebalanced(insertedKey int) *Node {
	balancedFactor := n.getBalancedFactor()
	println(fmt.Sprintf("rebalanced balancedFactor= %d,node.value= %d,insertedKey= %d", balancedFactor, n.Value, insertedKey))
	if balancedFactor > 1 && insertedKey < n.Value {
		if n.Left != nil && insertedKey > n.Left.Value {
			println("LR模式=> n.Left左旋, n右旋")
			n.Left = n.Left.leftRotate()
			return n.rightRotate()
		}
		println("LL模式=>n右旋")
		return n.rightRotate()
	}
	if balancedFactor < -1 && insertedKey > n.Value {
		if n.Right != nil && insertedKey < n.Right.Value {
			println("RL模式=>n.Right右旋,n左旋")
			n.Right = n.Right.rightRotate()
			return n.leftRotate()
		}
		println("RR模式=>n左旋")
		return n.leftRotate()
	}
	println("未失衡，什么也不干")
	return n
}

func (n *Node) getBalancedFactor() int {
	h1, h2 := 0, 0
	if n.Left != nil {
		h1 = n.Left.Height
	}
	if n.Right != nil {
		h2 = n.Right.Height
	}
	return h1 - h2
}

func (n *Node) leftRotate() *Node {
	y := n.Right
	y.Left, n.Right = n, y.Left
	n.reHeight()
	y.reHeight()
	return y
}

func (n *Node) rightRotate() *Node {
	x := n.Left
	x.Right, n.Left = n, n.Left.Right
	n.reHeight()
	x.reHeight()
	return x
}

func (n *Node) reHeight() int {
	n.Height = maxHeight(n.Left, n.Right) + 1

	return n.Height
}

func maxHeight(n1, n2 *Node) int {
	h1, h2 := 0, 0
	if n1 != nil {
		h1 = n1.reHeight()
	}
	if n2 != nil {
		h2 = n2.reHeight()
	}
	if h1 > h2 {
		return h1
	} else {
		return h2
	}
}

func (n *Node) compare(node *Node) int {
	return n.Value - node.Value
}

func (at *AvlTree) String() string {
	nodes := []*Node{at.root}
	values := make([][]int, 0)
	values = append(values, []int{at.root.Value})
	for len(nodes) > 0 {
		tmp := make([]*Node, 0)
		tmpValue := make([]int, 0)
		for _, item := range nodes {
			if item.Left != nil {
				tmp = append(tmp, item.Left)
				tmpValue = append(tmpValue, item.Left.Value)
			}
			if item.Right != nil {
				tmp = append(tmp, item.Right)
				tmpValue = append(tmpValue, item.Right.Value)
			}
		}
		nodes = tmp
		if len(tmpValue) != 0 {
			values = append(values, tmpValue)
		}
	}
	return fmt.Sprintf("%v", values)
}
