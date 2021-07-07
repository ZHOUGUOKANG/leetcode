package Num_451_FrequencySort

import "sort"

func frequencySort(s string) string {
	arr := []byte(s)
	mapping := make(map[byte]int)
	for _, item := range arr {
		if _, ok := mapping[item]; ok {
			mapping[item]++
		} else {
			mapping[item] = 1
		}
	}
	sort.Slice(arr, func(i, j int) bool {
		l, r := mapping[arr[i]], mapping[arr[j]]
		if l > r {
			return false
		}
		if l < r {
			return true
		}
		return arr[i] < arr[j]
	})
	return string(arr)
}
