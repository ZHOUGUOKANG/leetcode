package Num_5_LongestPalindrome

import "strings"

func LongestPalindrome(s string) string {
	str := []byte(s)
	length := len(str)
	if length < 2 {
		return s
	}
	tmpLen := length*2 + 1
	tmpStrArray := make([]byte, tmpLen)
	for i, j := 0, 0; i < tmpLen; i++ {
		if i%2 == 0 {
			//插入#号
			tmpStrArray[i] = 35
		} else {
			tmpStrArray[i] = str[j]
			j++
		}

	}
	max := ""
	tmpStr := ""
	index := 1
	for key, value := range tmpStrArray {
		tmpStr = string(value)
		index = 1
		for {
			left := key - index
			right := key + index
			if right >= tmpLen || left < 0 {
				if len(tmpStr) > len(max) {
					max = tmpStr
				}
				break
			}
			if tmpStrArray[left] == tmpStrArray[right] {
				tmpStr = string(tmpStrArray[left]) + tmpStr + string(tmpStrArray[left])
				index++
			} else {
				if len(tmpStr) > len(max) {
					max = tmpStr
				}
				break
			}
		}
	}
	return strings.Replace(max, "#", "", -1)
}
