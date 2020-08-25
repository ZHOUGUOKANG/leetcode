package Num_29_divide

import (
	"math"
)

func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}
	if divisor == 1 {
		return dividend
	}
	if divisor == -1 {
		if dividend == math.MinInt32 {
			return math.MaxInt32
		}
		return -dividend
	}
	sign := 1
	if dividend > 0 && divisor < 0 || dividend < 0 && divisor > 0 {
		sign = -1
	}
	if dividend > 0 {
		dividend = -dividend
	}
	if divisor > 0 {
		divisor = -divisor
	}
	if sign == -1 {
		return -div(dividend, divisor)
	} else {
		return div(dividend, divisor)
	}
}
func div(x, y int) (sum int) {
	if x <= y {
		tmp := y
		lastValue := y
		sum = 1
		for {
			lastValue = tmp
			tmp += tmp
			if tmp > x {
				sum += sum
			} else {
				sum += div(x-lastValue, y)
				break
			}
		}
	}
	return
}
