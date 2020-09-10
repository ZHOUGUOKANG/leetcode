package Num_49_groupAnagrams

import "sort"

func groupAnagrams(strs []string) [][]string {
	has := make(map[string]int)
	var ans [][]string
	for _, val := range strs {
		strSlice := []byte(val)
		sort.Slice(strSlice, func(i, j int) bool {
			return strSlice[i] > strSlice[j]
		})
		s := string(strSlice)
		if index, ok := has[s]; ok {
			ans[index] = append(ans[index], val)
		} else {
			ans = append(ans, []string{val})
			has[s] = len(ans) - 1
		}

	}
	return ans
}
