package Num_1_twoSum

//func twoSum(nums []int, target int) []int {
//	tmpNums := map[int]int{}
//	for key, value := range nums {
//		if index, ok := tmpNums[target-value]; ok {
//			return []int{index, key}
//		}
//		tmpNums[value] = key
//	}
//	return nil
//}

func twoSum(nums []int, target int) []int {
	n := len(nums)
	for l := 0; l < n-1; l++ {
		for r := n - 1; r > l; r-- {
			sum := nums[l] + nums[r]
			if sum == target {
				return []int{l, r}
			}
		}
	}
	return nil
}
