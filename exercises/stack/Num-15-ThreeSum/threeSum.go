package Num_15_ThreeSum

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	length := len(nums)
	result := make([][]int, 0)
	var start, end, sum int
	for current := 1; current < len(nums)-1; current++ {
		start, end = 0, length-1
		if current > 1 && nums[current] == nums[current-1] {
			start = current - 1
		}
		for start < current && current < end {
			if start > 0 && nums[start] == nums[start-1] {
				start++
				continue
			}

			if end < length-1 && nums[end] == nums[end+1] {
				end--
				continue
			}
			sum = nums[current] + nums[start] + nums[end]
			if sum < 0 {
				start++
			} else if sum > 0 {
				end--
			} else {
				result = append(result, []int{nums[current], nums[start], nums[end]})
				start++
				end--
			}
		}
	}
	return result
}
