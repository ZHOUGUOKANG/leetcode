package Num_31_nextPermutation

import (
	"sort"
)

/**
从后往前找第一组连续升序的两个数，然后从后往前找第一个比i-1大的数字并且交换，然后i后边的数组全部排序
*/
func nextPermutation(nums []int) {
	length := len(nums)
	for i := length - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			for j := length - 1; j > i-1; j-- {
				if nums[j] > nums[i-1] {
					nums[i-1], nums[j] = nums[j], nums[i-1]
					sort.Ints(nums[i:])
					return
				}
			}
		}
	}
	sort.Ints(nums)
}
