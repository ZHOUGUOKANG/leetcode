package Num_offer_09_Constructor

import "container/list"

type CQueue struct {
	stack1, stack2 *list.List
}

func Constructor() CQueue {
	return CQueue{
		stack1: list.New(),
		stack2: list.New(),
	}
}

func (c *CQueue) AppendTail(value int) {
	c.stack1.PushBack(value)
}

func (c *CQueue) DeleteHead() int {
	// 如果第二个栈为空
	if c.stack2.Len() == 0 {
		for c.stack1.Len() > 0 {
			c.stack2.PushBack(c.stack1.Remove(c.stack1.Back()))
		}
	}
	if c.stack2.Len() != 0 {
		e := c.stack2.Back()
		c.stack2.Remove(e)
		return e.Value.(int)
	}
	return -1
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
