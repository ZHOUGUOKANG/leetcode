package Num_3_LengthOflongestSubstring

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
