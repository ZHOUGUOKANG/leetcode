package Num_1370_SortString

func sortString(s string) string {
	words := make([]int, 26)
	for _, v := range s {
		words[v-'a']++
	}
	n := len(s)
	ans := make([]byte, n)
	for flag, cur := true, 0; cur < n; {
		for i := 0; i < 26; i++ {

			if flag {
				if words[i] > 0 {
					words[i]--
					ans[cur] = byte('a' + i)
					cur++
				}
			} else {
				if words[25-i] > 0 {
					words[25-i]--
					ans[cur] = byte('a' + 25 - i)
					cur++
				}
			}
		}
		flag = !flag
	}
	return string(ans)
}
