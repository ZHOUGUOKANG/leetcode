package Num_347_topKFrequent

import "container/heap"

func topKFrequent(nums []int, k int) []int {
	numsMap := make(map[int]int)
	for _, v := range nums {
		numsMap[v]++
	}
	h := &IHeap{}
	heap.Init(h)
	for key, value := range numsMap {
		heap.Push(h, &Item{key, value})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	ret := make([]int, k)
	for i := 0; i < k; i++ {
		ret[k-i-1] = heap.Pop(h).(*Item).key
	}
	return ret
}

type Item struct {
	key   int
	value int
}
type IHeap []*Item

func (h IHeap) Len() int           { return len(h) }
func (h IHeap) Less(i, j int) bool { return h[i].value < h[j].value }
func (h IHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IHeap) Push(x interface{}) {
	*h = append(*h, x.(*Item))
}

func (h *IHeap) Pop() interface{} {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[0 : n-1]
	return x
}
