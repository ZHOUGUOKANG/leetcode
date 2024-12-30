package num9

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x > 0 && x%10 == 0 {
		return false
	}
	var y int
	for t := x; t != 0; t = t / 10 {
		y = y*10 + t%10
		if x == y || t/10 == y {
			return true
		}
	}
	return false
}
