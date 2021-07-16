package Num_1818_MinAbsoluteSumDiff

import "sort"

func minAbsoluteSumDiff(nums1, nums2 []int) int {
	nums := append(sort.IntSlice{}, nums1...)
	nums.Sort()
	sum, maxN, n := 0, 0, len(nums)
	for i, v := range nums2 {
		diff := abs(nums1[i] - v)
		j := nums.Search(v)
		if j < n {
			maxN = max(maxN, diff-(nums[j]-v))
		}
		if j > 0 {
			maxN = max(maxN, diff-(v-nums[j-1]))
		}
		sum += diff
	}
	return (sum - maxN) % (1e9 + 7)
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
