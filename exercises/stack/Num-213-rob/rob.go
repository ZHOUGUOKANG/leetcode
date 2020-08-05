package Num_213_rob

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	s0, s1, ns0, ns1 := nums[0], nums[0], 0, nums[1]
	for i := 2; i < len(nums)-1; i++ {
		s0, ns0, s1, ns1 = max(s0, s1), max(ns0, ns1), s0+nums[i], ns0+nums[i]
	}
	return max(max(max(s1, ns1), max(s0, ns0)), ns0+nums[len(nums)-1])
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
