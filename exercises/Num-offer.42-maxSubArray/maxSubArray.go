package Num_offer_42_maxSubArray

func maxSubArray(nums []int) int {
	max, previous, current := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		current = nums[i]
		if previous > 0 {
			current = nums[i] + previous
		}
		if current > max {
			max = current
		}
		previous = current
	}
	return max
}
