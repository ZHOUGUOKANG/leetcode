package leetcodes

import (
	"math"
	"strconv"
)

/**
7. 整数反转
假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−231,  231 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0。
给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。
*/
func Reverse(x int) int {
	flag := true
	if x < 0 {
		x = -x
		flag = false
	}

	s1 := strconv.Itoa(x)
	s2 := []byte{}

	for i := len(s1) - 1; i >= 0; i-- {
		s2 = append(s2, s1[i])
	}

	x, _ = strconv.Atoi(string(s2))
	if !flag {
		x = -x
	}
	if x > math.MaxInt32 || x < math.MinInt32 {
		return 0
	} else {
		return x
	}
}
