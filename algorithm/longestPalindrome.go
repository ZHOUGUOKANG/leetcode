package leetcodes

import "strings"

/**
5. 最长回文子串
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

示例 1：

输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。
*/
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
