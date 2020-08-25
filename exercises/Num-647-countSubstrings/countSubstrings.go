package Num_647_countSubstrings

func countSubstrings(s string) int {
	count, n := 0, len(s)
	if n == 0 {
		return 0
	}
	for i := 0; i < n; i++ {
		count++
		if i > 0 && s[i-1] == s[i] {
			count++
			for j := 1; i-j-1 > -1 && i+j < n; j++ {
				if s[i-j-1] == s[i+j] {
					count++
				} else {
					break
				}
			}
		}
		if i > 0 && i < n-1 {
			for j := 1; (i-j > -1) && (i+j < n); j++ {
				if s[i-j] == s[i+j] {
					count++
				} else {
					break
				}
			}
		}
	}
	return count
}
