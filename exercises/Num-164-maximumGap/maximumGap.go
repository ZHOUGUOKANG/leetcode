package Num_164_MaximumGap

import "sort"

func maximumGap(nums []int) (ans int) {
	if len(nums) < 2 {
		return
	}
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if t := nums[i] - nums[i-1]; t > ans {
			ans = t
		}
	}
	return
}
