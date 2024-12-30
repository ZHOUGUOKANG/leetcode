package num8

import "math"

func myAtoi(s string) (ans int) {
	if len(s) == 0 {
		return 0
	}
	k := 1
	hasNum, hasSign, mustBreak := false, false, false
	for _, x := range s {
		switch {
		case x == ' ':
			if hasNum || hasSign {
				mustBreak = true
			}
		case x == '+':
			if hasNum || hasSign {
				mustBreak = true
				break
			}
			hasSign = true
		case x == '-':
			if hasNum || hasSign {
				mustBreak = true
				break
			}
			hasSign = true
			k = -1
		case x >= '0' && x <= '9':
			hasNum = true
			t := ans*10 + int(x-'0')
			if k == -1 && -t < math.MinInt32 {
				return math.MinInt32
			}
			if k == 1 && t > math.MaxInt32 {
				return math.MaxInt32
			}
			ans = t
		default:
			mustBreak = true
		}
		if mustBreak {
			break
		}
	}
	return k * ans
}
