package num12

import "bytes"

var lm = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
var vl = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

func intToRoman(num int) string {
	buf := bytes.NewBuffer([]byte{})
	for i, v := range vl {
		for ; num >= v; num -= v {
			buf.WriteString(lm[i])
		}
	}
	return buf.String()
}
