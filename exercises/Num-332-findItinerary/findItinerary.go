package Num_332_findItinerary

import "sort"

func findItinerary(tickets [][]string) []string {
	m := make(map[string][]string)
	var res []string
	for i := 0; i < len(tickets); i++ {
		from, to := tickets[i][0], tickets[i][1]
		m[from] = append(m[from], to)
	}
	for key := range m {
		sort.Strings(m[key])
	}
	var dfs func(curr string)
	dfs = func(curr string) {
		for {
			if v, ok := m[curr]; !ok || len(v) == 0 {
				break
			}
			tmp := m[curr][0]
			m[curr] = m[curr][1:]
			dfs(tmp)
		}
		res = append(res, curr)
	}
	dfs("JFK")
	n := len(res)
	for i := 0; i < n/2; i++ {
		res[i], res[n-1-i] = res[n-1-i], res[i]
	}
	return res
}
