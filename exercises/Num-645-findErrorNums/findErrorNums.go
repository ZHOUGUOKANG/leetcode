package Num_645_FindErrorNums

func findErrorNums(nums []int) []int {
	ans := make([]int, 2)
	for i := 0; i < len(nums); {
		j := nums[i] - 1
		if i == j || nums[i] == nums[j] {
			i++
			continue
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	for i, num := range nums {
		if num != i+1 {
			ans[0], ans[1] = num, i+1
			break
		}
	}
	return ans
}
