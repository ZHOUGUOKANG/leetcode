package Num_1021_removeOuterParentheses

import "strings"

func RemoveOuterParentheses(S string) string {
	//分解原语，规则：从左往右数，直到左括号和右括号数量相同
	var left, right int = 0, 0
	var tmp []byte
	var res strings.Builder
	for _, value := range S {
		if value == 40 {
			left++
		} else {
			right++
		}
		tmp = append(tmp, byte(value))
		if left == right {
			res.WriteString(string(tmp[1 : len(tmp)-1]))
			left, right = 0, 0
			tmp = []byte{}
		}
	}
	return res.String()
}
