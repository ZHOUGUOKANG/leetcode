package Num_167_twoSum

func twoSum(numbers []int, target int) []int {
	tmpNums := map[int]int{}
	for key, value := range numbers {
		if index, ok := tmpNums[target-value]; ok {
			return []int{index + 1, key + 1}
		}
		tmpNums[value] = key
	}
	return nil
}
