package Num_9_IsPalindrome

import "strconv"

func IsPalindrome(x int) bool {
	str := strconv.Itoa(x)
	length := len(str)
	if length < 2 {
		return true
	}
	for i := 0; i < length/2; i++ {
		//println(i,length-i-1,str[i],str[length-i-1])
		if str[i] != str[length-i-1] {
			return false
		}
	}
	return true
}
