package Num_06_03_findMagicIndex

func findMagicIndex(nums []int) int {
	for k, v := range nums {
		if k == v {
			return k
		}
	}
	return -1
}
