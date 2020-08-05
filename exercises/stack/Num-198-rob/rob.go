package Num_198_rob

func rob(nums []int) int {
	selected, noSelected := 0, 0
	for i := 0; i < len(nums); i++ {
		selected, noSelected = noSelected+nums[i], max(selected, noSelected)
	}
	return max(selected, noSelected)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
