package Num_53_Search

func search(nums []int, target int) int {
	n := len(nums)
	if n <= 0 || nums[n-1] < target || nums[0] > target {
		return 0
	}
	if nums[0] == target && nums[n-1] == target {
		return n
	}
	mid := n / 2
	return search(nums[:mid], target) + search(nums[mid:], target)
}
