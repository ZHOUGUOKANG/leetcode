package Num_12_IntToRoman

import (
	"leetcode/exercises"
)

func IntToRoman(num int) string {
	output := ""
	for _, value := range exercises.RomanSort {
		merchant := num / value
		if merchant > 0 {
			for i := 0; i < merchant; i++ {
				output += exercises.RomanNumeralDict[value]
			}
			num = num - merchant*value
		}
	}
	return output
}
