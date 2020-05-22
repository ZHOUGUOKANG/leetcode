package chart

import "strings"

func ShortestPath(relationMap map[string]map[string]int, start, end string) (string, int) {
	maps := [][]string{[]string{start}}
	expected := maps
	flag := true
	endResult := make([][]string, 0)
	for {
		if len(expected) == 0 {
			break
		}
		maps, expected = expected, make([][]string, 0)
		for _, item := range maps {

			result := make([]string, 0)
			for name, _ := range relationMap[item[len(item)-1]] {
				flag = true
				for _, key := range item {
					if name == key {
						flag = false
						break
					}
				}
				if flag {
					result = append(result, name)
				}
			}
			for _, r := range result {
				if r == end {
					endResult = append(endResult, append(append([]string{}, item...), r))
					continue
				} else {
					expected = append(expected, append(append([]string{}, item...), r))
				}
			}
		}
	}

	lessScore := -1
	lessPath := make([]string, 0)
	for _, item := range endResult {
		score := 0
		for i := 1; i < len(item); i++ {
			score += relationMap[item[i]][item[i-1]]
		}
		if lessScore == -1 {
			lessScore = score
			lessPath = item
		} else if score < lessScore {
			lessScore = score
			lessPath = item
		}
	}
	return strings.Join(lessPath, "-"), lessScore
}
