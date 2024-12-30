package num26

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	i, j := 0, 1
	for ; j < len(nums); j++ {
		if nums[j] == nums[i] {
			continue
		}
		i++
		nums[i] = nums[j]
	}
	return i + 1
}
