package Num_30_findSubstring

func findSubstring(s string, words []string) []int {
	w, l := len(words), len(s)
	var r []int
	if w == 0 || l == 0 {
		return r
	}
	n := len(words[0])
	if l < n*w {
		return r
	}
	wordsMap := make(map[string]int)
	for _, v := range words {
		wordsMap[v]++
	}
	for i := 0; i <= l-n*w; i++ {
		if _, ok := wordsMap[s[i:i+n]]; ok {
			ms := make(map[string]int, len(wordsMap))
			for c, v := range wordsMap {
				ms[c] = v
			}
			flag := true
			for j := 0; j < w; j++ {
				t := i + n*j
				if v, ok := ms[s[t:t+n]]; ok && v > 0 {
					ms[s[t:t+n]]--
				} else {
					flag = false
					break
				}
			}
			if flag {
				r = append(r, i)
			}
		}
	}
	return r
}
