package Num_459_repeatedSubstringPattern

func repeatedSubstringPattern(s string) bool {
	n := len(s)
	for i := n / 2; i >= 1; i-- {
		if n%i != 0 {
			continue
		}
		flag := true
		for j := i; j < n; j++ {
			t := j % i
			if s[t] != s[j] {
				flag = false
				break
			}
		}
		if flag {
			return flag
		}
	}
	return false
}
