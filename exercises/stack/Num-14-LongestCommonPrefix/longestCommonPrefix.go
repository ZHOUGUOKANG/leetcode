package Num_14_LongestCommonPrefix

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	minLen := 0
	for key, value := range strs {
		if key == 0 {
			minLen = len(value)
			continue
		}
		if len(value) < minLen {
			minLen = len(value)
		}
	}
	tmpStr := ""
	flag := true
	for i := 0; i < minLen && flag; i++ {
		tmpPre := ""
		for key, value := range strs {
			if key == 0 {
				tmpPre = string(value[i])
				continue
			}
			if tmpPre != string(value[i]) {
				flag = false
				break
			}
			if key == len(strs)-1 && tmpPre == string(value[i]) {
				tmpStr += tmpPre
			}
		}
	}
	return tmpStr
}
