package list

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

var errorNotFound = errors.New("value not found")

type Node struct {
	Value   int
	forward []*Node
}

type SkipList struct {
	head     *Node
	MaxLevel int
	Level    int
}

func NewSkipList(maxLevel int) *SkipList {
	return &SkipList{MaxLevel: maxLevel, head: newNode(0, maxLevel), Level: 1}
}
func newNode(value, level int) *Node {
	return &Node{Value: value, forward: make([]*Node, level)}
}

const Probability = 0.25 // 基于时间与空间综合 best practice 值, 越上层概率越小

func (list *SkipList) randLevel() (level int) {
	rand.Seed(time.Now().UnixNano())
	for level = 1; rand.Float32() < Probability && level < list.MaxLevel; level++ {
	}
	return
}
func (list *SkipList) Insert(value int) *Node {
	curr := list.head
	update := make([]*Node, list.MaxLevel)
	for i := list.Level - 1; i >= 0; i-- {
		if curr.forward[i] == nil || curr.forward[i].Value > value {
			update[i] = curr
		} else {
			for curr.forward[i] != nil && curr.forward[i].Value < value {
				curr = curr.forward[i]
			}
			update[i] = curr
		}
	}
	level := list.randLevel()
	if level > list.Level {
		for i := list.Level; i < level; i++ {
			update[i] = list.head
		}
		list.Level = level
	}
	node := newNode(value, list.Level)
	for i := 0; i < level; i++ {
		node.forward[i] = update[i].forward[i]
		update[i].forward[i] = node
	}

	return node
}
func (list *SkipList) Delete(value int) (node *Node) {
	current := list.head
	for i := list.Level - 1; i >= 0; i-- {
		for current.forward[i] != nil {
			if current.forward[i].Value == value {
				node = current.forward[i]
				current.forward[i] = node.forward[i]
				node.forward[i] = nil
			} else if current.forward[i].Value > value {
				break
			} else {
				current = current.forward[i]
			}
		}
	}
	return node
}
func (list *SkipList) Search(value int) (node *Node, err error) {
	current := list.head
	for i := list.Level - 1; i >= 0; i-- {
		for current.forward[i] != nil {
			if current.forward[i].Value == value {
				node = current.forward[i]
				break
			} else if current.forward[i].Value > value {
				break
			} else {
				current = current.forward[i]
			}
		}
	}
	if node == nil {
		return nil, errorNotFound
	}
	return node, nil
}
func (list *SkipList) Print() {
	fmt.Println()
	for i := list.Level - 1; i >= 0; i-- {
		current := list.head
		for current.forward[i] != nil {
			fmt.Printf("%d ", current.forward[i].Value)
			current = current.forward[i]
		}
		fmt.Printf("***************** Level %d \n", i+1)
	}
}
