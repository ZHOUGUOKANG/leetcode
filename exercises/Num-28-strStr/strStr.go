package Num_28_strStr

func strStr(haystack string, needle string) int {
	h, n := len(haystack), len(needle)
	if n == 0 {
		return 0
	}
	if h == 0 {
		return -1
	}
	for i := 0; i <= h-n; i++ {
		if haystack[i] == needle[0] {
			flag := true
			for x, y := i+1, 1; y < n; x, y = x+1, y+1 {
				if haystack[x] != needle[y] {
					flag = false
					break
				}
			}
			if flag {
				return i
			}
		}
	}
	return -1
}
