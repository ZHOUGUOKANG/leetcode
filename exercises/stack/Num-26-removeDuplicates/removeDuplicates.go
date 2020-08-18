package Num_26_removeDuplicates

func removeDuplicates(nums []int) int {
	if len(nums) == 0 || len(nums) == 1 {
		return len(nums)
	}
	left, right := 0, 1
	for right < len(nums) {
		if nums[right] != nums[left] {
			left++
			nums[left] = nums[right]
		}
		right++
	}
	return left + 1
}
