package Num_210_findOrder

func findOrder(numCourses int, prerequisites [][]int) []int {
	chart := make([][]int, numCourses)
	visited := make([]int, numCourses)
	stack := make([]int, 0)
	flag := true
	dfs := func(u int) {}
	for i := 0; i < len(prerequisites); i++ {
		chart[prerequisites[i][0]] = append(chart[prerequisites[i][0]], prerequisites[i][1])
	}
	dfs = func(u int) {
		visited[u] = 1
		for i := 0; i < len(chart[u]); i++ {
			if visited[chart[u][i]] == 0 {
				dfs(chart[u][i])
			} else if visited[chart[u][i]] == 1 {
				flag = false
				return
			}
		}
		visited[u] = 2
		stack = append(stack, u)
	}
	for i := 0; i < numCourses && flag; i++ {
		if visited[i] == 0 {
			dfs(i)
		}
	}
	if flag {
		return stack
	} else {
		return []int{}
	}
}
