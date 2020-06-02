package Num_1218_longestSubsequence

func longestSubsequence(arr []int, difference int) int {
	if len(arr) < 2 {
		return len(arr)
	}
	result := make(map[int]int, len(arr))
	max := 1
	for _, value := range arr {
		if r, ok := result[value-difference]; ok {
			result[value] = r + 1
		} else {
			result[value] = 1
		}
		if result[value] > max {
			max = result[value]
		}
	}
	return max
}
