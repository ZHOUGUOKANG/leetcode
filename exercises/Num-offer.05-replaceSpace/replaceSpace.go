package Num_offer_05_replaceSpace

func replaceSpace(s string) string {
	n := len(s)
	count := 0
	for i := 0; i < n; i++ {
		if s[i] == ' ' {
			count++
		}
	}
	if count == 0 {
		return s
	}
	ret := make([]byte, n+(count*2))
	for i, j := 0, 0; i < n; i++ {
		if s[i] == ' ' {
			ret[j], ret[j+1], ret[j+2] = '%', '2', '0'
			j += 3
		} else {
			ret[j] = s[i]
			j++
		}
	}
	return string(ret)
}
