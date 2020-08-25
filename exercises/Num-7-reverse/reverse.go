package Num_7_reverse

import (
	"math"
	"strconv"
)

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
