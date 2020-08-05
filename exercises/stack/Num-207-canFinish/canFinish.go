package Num_207_canFinish

func canFinish(numCourses int, prerequisites [][]int) bool {
	chart := make([][]int, numCourses)
	visited := make([]int, numCourses)
	valid := true
	dfs := func(courseID int) {}
	for i := 0; i < len(prerequisites); i++ {
		chart[prerequisites[i][1]] = append(chart[prerequisites[i][1]], prerequisites[i][0])
	}
	dfs = func(courseID int) {
		visited[courseID] = 1
		for _, v := range chart[courseID] {
			if visited[v] == 0 {
				dfs(v)
			} else if visited[v] == 1 {
				valid = false
				return
			}
		}
		for i := 0; i < len(chart[courseID]); i++ {
			if visited[chart[courseID][i]] == 0 {
				dfs(chart[courseID][i])
			} else if visited[chart[courseID][i]] == 1 {
				valid = false
				return
			}
		}
		visited[courseID] = 2
	}
	for i := 0; i < numCourses && valid; i++ {
		if visited[i] == 0 {
			dfs(i)
		}
	}
	return valid
}
