package num30

func findSubstring(s string, words []string) (ans []int) {
	wordsMap := make(map[string]int, len(words))
	for _, word := range words {
		wordsMap[word]++
	}
	ls, m, n := len(s), len(words), len(words[0])
	curWordsMap := make(map[string]int, len(words))
	for i := 0; i < n; i++ {
		l, r, k := i, i, 0
		clear(curWordsMap)
		for r+n <= ls {
			w := s[r : r+n]
			k++
			r += n
			curWordsMap[w]++
			for curWordsMap[w] > wordsMap[w] {
				curWordsMap[s[l:l+n]]--
				l += n
				k--
			}
			if k == m {
				ans = append(ans, l)
			}
		}
	}
	return ans
}

//func findSubstring(s string, words []string) (ans []int) {
//	wordsMap := make(map[string]int, len(words))
//	ls, m, n := len(s), len(words), len(words[0])
//	for i := 0; i <= ls-m*n; i++ {
//		clear(wordsMap)
//		for _, word := range words {
//			wordsMap[word]++
//		}
//		flag := true
//		for j := i; j < i+(m-1)*n; j += n {
//			if _, ok := wordsMap[s[j:j+n]]; !ok {
//				flag = false
//				break
//			}
//			wordsMap[s[j:j+n]]--
//			if wordsMap[s[j:j+n]] < 0 {
//				flag = false
//				break
//			}
//		}
//		if flag {
//			ans = append(ans, i)
//		}
//	}
//	return ans
//}
