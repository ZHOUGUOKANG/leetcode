package Num_214_shortestPalindrome

//func shortestPalindrome(s string) string {
//	if len(s) == 0 {
//		return s
//	}
//	r := []rune(s)
//	if isPalindrome(r,0){
//		return s
//	}
//	reverse(r)
//	i := 1
//	tmp := make([]rune,len(r))
//	for {
//		copy(tmp[:i],r[:i])
//		reverse(tmp[:i])
//		r = append(r,tmp[:i]...)
//		if isPalindrome(r,i){
//			break
//		}else {
//			r = r[:len(r)-i]
//			i++
//		}
//	}
//	return string(r)
//}
//
//func reverse(r []rune)  {
//	n:=len(r)
//	for i:=0;i<n/2;i++{
//		r[i],r[n-i-1] = r[n-i-1],r[i]
//	}
//}
//
//func isPalindrome(s []rune,i int) bool {
//	n := len(s)
//	for j := n/2;j >= i;j--{
//		if s[j] != s[n-j-1]{
//			return false
//		}
//	}
//	return true
//}

func shortestPalindrome(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}
	r := []rune(s)
	reverse(r)
	for left, right := 0, n; left < n; left++ {
		if isPalindrome(r[left:right]) {
			return string(r) + s[n-left:]
		}
	}
	return string(r) + s[1:]
}
func reverse(r []rune) {
	n := len(r)
	for i := 0; i < n/2; i++ {
		r[i], r[n-i-1] = r[n-i-1], r[i]
	}
}

func isPalindrome(s []rune) bool {
	n := len(s)
	for l, r := 0, n-1; l < r; l, r = l+1, r-1 {
		if s[l] != s[r] {
			return false
		}
	}
	return true
}
