package Num_18_fourSum

import "sort"

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	return nSum(nums, target, 4)
}

func nSum(nums []int, target int, n int) [][]int {
	l := len(nums)
	if l < n {
		return nil
	}
	if n == 1 {
		for i := 0; i < l; i++ {
			if nums[i] == target {
				return [][]int{{nums[i]}}
			}
		}
		return nil
	}
	var result [][]int
	for i := 0; i <= l-n; i++ {
		val := nums[i]
		if i > 0 && val == nums[i-1] {
			continue
		}
		r := nSum(nums[i+1:], target-val, n-1)
		if r == nil {
			continue
		}
		for _, v := range r {
			v = append(v, val)
			result = append(result, v)
		}
	}
	return result
}
