package Num_976_LargestPerimeter

import "sort"

func largestPerimeter(A []int) int {
	n := len(A)
	sort.Ints(A)
	for i := n - 1; i >= 2; i-- {
		if c := A[i-1] + A[i-2]; c > A[i] {
			return c + A[i]
		}
	}
	return 0
}
