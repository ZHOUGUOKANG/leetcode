package num4

import "math"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	var find func(i, j, k int) int
	find = func(i, j, k int) int {
		if i >= m {
			return nums2[j+k-1]
		}
		if j >= n {
			return nums1[i+k-1]
		}
		if k == 1 {
			return min(nums1[i], nums2[j])
		}
		p := k / 2
		x, y := math.MaxInt64, math.MaxInt64
		if n1 := i + p - 1; n1 < m {
			x = nums1[n1]
		}
		if n2 := j + p - 1; n2 < n {
			y = nums2[n2]
		}
		if x < y {
			return find(i+p, j, k-p)
		}
		return find(i, j+p, k-p)
	}
	n1, n2 := find(0, 0, (m+n+1)/2), find(0, 0, (m+n+2)/2)
	return float64(n1+n2) / 2
}
