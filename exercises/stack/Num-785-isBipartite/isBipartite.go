package Num_785_isBipartite

func isBipartite(graph [][]int) bool {
	var result = make([]int, len(graph))
	RED := 1
	//BLUE := -1
	result[0] = RED
	for i := 0; i < len(graph); i++ {
		if result[i] == 0 {
			continue
		}

		next := -result[i]
		for _, node := range graph[i] {
			if result[node] == 0 {
				result[node] = next
			} else {
				if result[node] == result[i] {
					return false
				}
			}
		}
	}
	for key, value := range result {
		if value == 0 {
			for _, v := range graph[key] {
				if result[v] != 0 {
					result[key] = -result[v]
					break
				}
			}
			if result[key] == 0 {
				result[key] = RED
			}
			for _, v := range graph[key] {
				if result[key] == result[v] {
					return false
				}
			}
		}
	}
	return true
}
