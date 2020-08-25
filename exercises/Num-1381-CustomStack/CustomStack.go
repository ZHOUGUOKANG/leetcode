package Num_1381_CustomStack

type CustomStack struct {
	list    []int
	maxSize int
}

func Constructor(maxSize int) CustomStack {
	list := make([]int, 0)
	return CustomStack{list, maxSize}
}

func (this *CustomStack) Push(x int) {
	if len(this.list) < this.maxSize {
		this.list = append(this.list, x)
	}
}

func (this *CustomStack) Pop() int {
	top := -1
	if len(this.list) > 0 {
		top = this.list[len(this.list)-1]
		this.list = this.list[:len(this.list)-1]
	}
	return top
}

func (this *CustomStack) Increment(k int, val int) {
	if len(this.list) < k {
		k = len(this.list)
	}
	for i := 0; i < k; i++ {
		this.list[i] += val
	}
}

func (this *CustomStack) GetList() []int {
	return this.list
}

/**
 * Your CustomStack object will be instantiated and called as such:
 * obj := Constructor(maxSize);
 * obj.Push(x);
 * param_2 := obj.Pop();
 * obj.Increment(k,val);
 */
