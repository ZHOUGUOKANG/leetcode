package Num_33_search

func search(nums []int, target int) int {
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		midNum := nums[mid]
		if midNum == target {
			return mid
		}
		if midNum >= nums[0] {
			if target >= nums[left] && target < midNum {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if target > midNum && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}
