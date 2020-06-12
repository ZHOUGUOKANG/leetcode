package Num_8_MyAtoi

import (
	"math"
)

func MyAtoi(str string) int {
	var num int64 = 0
	isPositiveNumber := true
	hasSymbol := false
	start := false
	for _, value := range str {
		//空格
		if value == 32 {
			if start || hasSymbol {
				break
			} else {
				continue
			}
		}
		if value < 43 || value > 57 || value == 44 || value == 46 || value == 47 {
			break
		}
		//43是 + 号
		if value == 43 {
			if hasSymbol {
				break
			}
			hasSymbol = true
			continue
		}
		//45是 - 号
		if value == 45 {
			if hasSymbol {
				break
			}
			hasSymbol = true
			isPositiveNumber = false
			continue
		}
		if num > math.MaxInt32 || num < math.MinInt32 {
			break
		}
		hasSymbol = true
		start = true
		num = num*10 + int64(value) - 48
	}
	if !isPositiveNumber {
		num = -num
	}
	if num > math.MaxInt32 {
		return math.MaxInt32
	}
	if num < math.MinInt32 {
		return math.MinInt32
	}
	return int(num)
}
