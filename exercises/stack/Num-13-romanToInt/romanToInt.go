package Num_13_romanToInt

import "leetcode/exercises"

func RomanToInt(s string) int {
	num := 0
	length := len(s)
	tmpStr := ""
	for i := 0; i < length; i++ {
		tmpStr += string(s[i])
		if i < length-1 {
			if _, ok := exercises.NumeralRomanDict[tmpStr+string(s[i+1])]; ok {
				i++
				tmpStr += string(s[i])
			}
		}
		num += exercises.NumeralRomanDict[tmpStr]
		tmpStr = ""
	}
	return num
}
