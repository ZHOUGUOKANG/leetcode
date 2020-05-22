package Num_394_DecodeString

import "strconv"

func decodeString(s string) string {
	left := false
	leftNum := 0
	stack := make([]int32, 0)
	tmpStack := make([]int32, 0)
	tmpNum := ""
	for _, value := range s {
		switch {
		// [ 不断压入栈
		case value == 91:
			if left {
				leftNum++
				tmpStack = append(tmpStack, value)
				continue
			}
			left = true
			leftNum++
			continue
		case value == 93:
			if left {
				leftNum--
				if leftNum > 0 {
					tmpStack = append(tmpStack, value)
					continue
				}
			}
			tmpStr := []int32(decodeString(string(tmpStack)))
			num, _ := strconv.Atoi(tmpNum)
			for i := 0; i < num; i++ {
				stack = append(stack, tmpStr...)
			}
			tmpNum = ""
			tmpStack = tmpStack[:0]
			left = false
			continue
			//记录数字是多少
		case isNumber(value):
			if left {
				tmpStack = append(tmpStack, value)
				continue
			}
			tmpNum += string(byte(value))
			continue
		default:
			if left {
				tmpStack = append(tmpStack, value)
			} else {
				stack = append(stack, value)
			}
			//入栈
		}
	}
	return string(stack)
}

func isNumber(byteNum int32) bool {
	if byteNum > 47 && byteNum < 58 {
		return true
	}
	return false
}
