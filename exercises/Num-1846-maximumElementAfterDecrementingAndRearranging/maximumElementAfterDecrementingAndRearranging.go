package Num_1846_MaximumElementAfterDecrementingAndRearranging

import "sort"

func maximumElementAfterDecrementingAndRearranging(arr []int) int {
	sort.Ints(arr)
	arr[0] = 1
	for i := 1; i < len(arr); i++ {
		arr[i] = min(arr[i], arr[i-1]+1)
	}
	return arr[len(arr)-1]
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
