package Num_32_longestValidParentheses

func longestValidParentheses(s string) int {
	maxAns := 0
	var result []int
	result = append(result, -1)
	for i := 0; i < len(s); i++ {
		//（
		if s[i] == 40 {
			result = append(result, i)
		} else {
			result = result[:len(result)-1]
			if len(result) == 0 {
				result = append(result, i)
			} else {
				maxAns = max(maxAns, i-result[len(result)-1])
			}
		}
	}
	return maxAns
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
