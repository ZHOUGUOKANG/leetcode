package num27

func removeElement(nums []int, val int) int {
	i, j := 0, 0
	for ; i < len(nums) && j < len(nums); i++ {
		if nums[i] == val {
			for j = max(j, i); j < len(nums); j++ {
				if nums[j] == val {
					continue
				}
				nums[i], nums[j] = nums[j], nums[i]
				break
			}
			if j >= len(nums) {
				return i
			}
		}
	}
	return i
}
