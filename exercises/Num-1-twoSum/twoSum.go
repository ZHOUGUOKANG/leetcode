package Num_1_twoSum

func TwoSum(nums []int, target int) []int {
	tmpNums := map[int]int{}
	for key, value := range nums {
		if index, ok := tmpNums[target-value]; ok {
			return []int{index, key}
		}
		tmpNums[value] = key
	}
	return nil
}
