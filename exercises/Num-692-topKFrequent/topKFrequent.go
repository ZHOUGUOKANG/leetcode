package Num_692_TopKFrequent

import (
	"container/heap"
)

type ret []string

func (r *ret) Len() int {
	return len(*r)
}

func (r *ret) Less(i, j int) bool {
	x, y := (*r)[i], (*r)[j]
	xVal, yVal := mapping[x], mapping[y]
	return xVal < yVal || xVal == yVal && x > y
}

func (r *ret) Swap(i, j int) {
	(*r)[i], (*r)[j] = (*r)[j], (*r)[i]
}

func (r *ret) Push(x interface{}) {
	*r = append(*r, x.(string))
}

func (r *ret) Pop() interface{} {
	val := (*r)[r.Len()-1]
	*r = (*r)[:r.Len()-1]
	return val
}

var mapping map[string]int

func topKFrequent(words []string, k int) []string {
	mapping = make(map[string]int)
	for _, w := range words {
		mapping[w]++
	}
	result := &ret{}
	for word, _ := range mapping {
		heap.Push(result, word)
		if result.Len() > k {
			heap.Pop(result)
		}
	}
	r := make([]string, k)
	for i := k - 1; i >= 0; i-- {
		r[i] = heap.Pop(result).(string)
	}
	return r
}
