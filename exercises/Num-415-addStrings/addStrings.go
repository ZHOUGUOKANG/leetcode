package Num_415_addStrings

//48 is 0
//func addStrings(num1 string, num2 string) string {
//	m,n := len(num1)-1,len(num2)-1
//	result := ""
//	var left,right,carry,cha,sum uint8 = 0,0,0,0,0
//	for  {
//		if m >= 0{
//			left = num1[m] - 48
//		}else {
//			left = 0
//		}
//		if n >= 0{
//			right = num2[n] - 48
//		}else {
//			right = 0
//		}
//		sum = left + right + carry
//		carry = sum / 10
//		cha = sum - carry * 10
//		result = string(cha + 48) + result
//		m--
//		n--
//		if m < 0 && n < 0{
//			if carry != 0{
//				result = string(carry + 48) + result
//			}
//			break
//		}
//	}
//	return result
//}

func addStrings(num1 string, num2 string) string {
	if len(num1) < len(num2) {
		num1, num2 = num2, num1
	}
	maxLen := len(num1)
	result := make([]byte, maxLen+1)
	var left, right, carry, cha, sum uint8 = 0, 0, 0, 0, 0
	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0 || carry > 0; i, j = i-1, j-1 {
		if i < 0 {
			left = 48
		} else {
			left = num1[i]
		}
		if j < 0 {
			right = 48
		} else {
			right = num2[j]
		}
		sum = left + right + carry - 96
		if sum >= 10 {
			carry = 1
		} else {
			carry = 0
		}
		cha = sum - carry*10
		result[maxLen] = cha + 48
		maxLen--
	}
	if result[0] == 0 {
		result = result[1:]
	}
	return string(result)
}
