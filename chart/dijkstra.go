package chart

import (
	"strings"
)

func Dijkstra(g map[string]map[string]int, start, end string) (string, int) {
	costs := map[string]int{start: 0}
	target, lastTarget := start, start
	paths := []string{start}
	for target != end {
		tmpV, tmpK := -1, ""
		for k, v := range g[target] {
			if k == start || k == lastTarget {
				continue
			}
			if tmpV == -1 || v < tmpV {
				tmpV, tmpK = v, k
			}
			if value, ok := costs[k]; !ok || value < v {
				costs[k] = costs[target] + v
			}
		}
		target, lastTarget = tmpK, target
		paths = append(paths, target)
	}
	return strings.Join(paths, "-"), costs[end]
}
