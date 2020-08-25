package Num_336_palindromePairs

func palindromePairs(words []string) [][]int {
	wordsRev := make([]string, len(words))
	indices := map[string]int{}
	for i := 0; i < len(words); i++ {
		w := reverse(words[i])
		wordsRev[i] = w
		indices[w] = i
	}
	var ret [][]int
	for i := 0; i < len(words); i++ {
		word := words[i]
		m := len(word)
		if m == 0 {
			continue
		}
		for j := 0; j <= m; j++ {
			if isPalindrome(word, j, m-1) {
				leftId := findWord(word, 0, j-1, indices)
				if leftId != -1 && leftId != i {
					ret = append(ret, []int{i, leftId})
				}
			}
			if j != 0 && isPalindrome(word, 0, j-1) {
				rightId := findWord(word, j, m-1, indices)
				if rightId != -1 && rightId != i {
					ret = append(ret, []int{rightId, i})
				}
			}
		}
	}
	return ret
}

func findWord(s string, left, right int, indices map[string]int) int {
	if v, ok := indices[s[left:right+1]]; ok {
		return v
	}
	return -1
}

func isPalindrome(s string, left, right int) bool {
	for i := 0; i < (right-left+1)/2; i++ {
		if s[left+i] != s[right-i] {
			return false
		}
	}
	return true
}

func reverse(s string) string {
	n := len(s)
	b := []byte(s)
	for i := 0; i < n/2; i++ {
		b[i], b[n-i-1] = b[n-i-1], b[i]
	}
	return string(b)
}

//func palindromePairs(words []string) [][]int {
//	caches := make([][]byte,len(words))
//	for i:=0;i<len(words);i++{
//		caches[i] = []byte(words[i])
//	}
//	result := make([][]int,0)
//	for i:=0;i<len(words);i++{
//		for j:=i+1;j<len(words);j++{
//			if isPalindrome(caches[i],caches[j]){
//				result = append(result,[]int{i,j})
//			}
//			if isPalindrome(caches[j],caches[i]){
//				result = append(result,[]int{j,i})
//			}
//		}
//	}
//	return result
//}
//
//func isPalindrome(s1,s2 []byte)bool{
//	s1 = append(s1,s2...)
//	length := len(s1)
//	for i:=0;i<length/2;i++{
//		if s1[i] != s1[length - i - 1]{
//			return false
//		}
//	}
//	return true
//}
