package Num_35_searchInsert

func searchInsert(nums []int, target int) int {
	length := len(nums)
	left, right := 0, length-1
	current := length
	mid := 0
	for left <= right {
		mid = (right-left)>>1 + left
		if nums[mid] >= target {
			current = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return current
}

/**
func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	for i := 0; i < len(nums); i++ {
		if target <= nums[i] {
			return i
		}
	}
	return len(nums)
}
*/
