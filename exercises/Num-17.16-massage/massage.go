package Num_17_16_massage

func massage(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	x, y, z := 0, 0, nums[0]
	for i := 1; i < len(nums); i++ {
		x, y, z = max(x, y), z, max(x, y)+nums[i]
	}
	println(x, y, z)
	return max(y, z)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
