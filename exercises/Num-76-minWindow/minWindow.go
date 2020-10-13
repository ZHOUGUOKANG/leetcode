package Num_76_MinWindow

import "math"

func minWindow(s string, t string) string {
	ori, cur := make(map[byte]int), make(map[byte]int)
	for i := 0; i < len(t); i++ {
		ori[t[i]]++
	}
	check := func() bool {
		for k, v := range ori {
			if val, ok := cur[k]; !ok || val < v {
				return false
			}
		}
		return true
	}
	tmpLen, tmpL, tmpR := math.MaxInt32, 0, 0
	for l, r := 0, 0; r < len(s); r++ {
		if _, ok := ori[s[r]]; ok {
			cur[s[r]]++
		}
		for check() && l <= r {
			if r-l < tmpLen {
				tmpLen, tmpL, tmpR = r-l, l, r
			}
			if _, ok := cur[s[l]]; ok {
				cur[s[l]]--
			}
			l++
		}
	}
	if tmpLen == math.MaxInt32 {
		return ""
	}
	return s[tmpL : tmpR+1]
}
