package leetcodes

/**
3. 无重复字符的最长子串
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

滑动窗口法
*/
func LengthOfLongestSubstring(s string) int {
	b := []byte(s)
	lenStr := len(b)
	result := make(map[byte]int, lenStr)
	maxStrLen := 0

	tmpLen := 0
	tmpMin := 0
	for min, max := 0, 0; max < lenStr; max++ {
		if index, ok := result[b[max]]; ok {
			if b[min] == b[index] {
				min++
			} else {
				tmpMin = result[b[max]] + 1
				for i := 0; i < result[b[max]]-min+1; i++ {
					delete(result, b[min+i])
				}
				min = tmpMin
			}
		}
		result[b[max]] = max
		tmpLen = max - min + 1
		if tmpLen > maxStrLen {
			maxStrLen = tmpLen
		}
	}
	return maxStrLen
}
