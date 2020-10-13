package Num_834_SumOfDistancesInTree

//func sumOfDistancesInTree(n int, edges [][]int) []int {
//	graph := make([][]int, n)
//	for _, e := range edges {
//		u, v := e[0], e[1]
//		graph[u] = append(graph[u], v)
//		graph[v] = append(graph[v], u)
//	}
//
//	sz := make([]int, n)
//	dp := make([]int, n)
//	var dfs func(u, f int)
//	dfs = func(u, f int) {
//		sz[u] = 1
//		for _, v := range graph[u] {
//			if v == f {
//				continue
//			}
//			dfs(v, u)
//			dp[u] += dp[v] + sz[v]
//			sz[u] += sz[v]
//		}
//	}
//	dfs(0, -1)
//	//println(fmt.Sprintf("%v %v",dp,sz))
//
//	ans := make([]int, n)
//	var dfs2 func(u, f int)
//	dfs2 = func(u, f int) {
//		ans[u] = dp[u]
//		for _, v := range graph[u] {
//			if v == f {
//				continue
//			}
//			pu, pv := dp[u], dp[v]
//			su, sv := sz[u], sz[v]
//
//			dp[u] -= dp[v] + sz[v]
//			sz[u] -= sz[v]
//			dp[v] += dp[u] + sz[u]
//			sz[v] += sz[u]
//
//			dfs2(v, u)
//
//			dp[u], dp[v] = pu, pv
//			sz[u], sz[v] = su, sv
//		}
//	}
//	dfs2(0, -1)
//	return ans
//}
func sumOfDistancesInTree(N int, edges [][]int) []int {
	childNums := make([]int, N)
	ans := make([]int, N)
	tMap := make([][]int, N)
	for i := 0; i < N; i++ {
		childNums[i] = 1
	}
	for _, v := range edges {
		f, t := v[0], v[1]
		tMap[f] = append(tMap[f], t)
		tMap[t] = append(tMap[t], f)
	}
	var dfsRoot func(cur, root int)
	dfsRoot = func(cur, root int) {
		for _, next := range tMap[cur] {
			if next == root {
				continue
			}
			dfsRoot(next, cur)
			//计算子节点数量
			childNums[cur] += childNums[next]
			//路径和=所有子节点的路径和+子节点数量
			ans[cur] += ans[next] + childNums[next]
		}
	}
	dfsRoot(0, -1)
	var dfsAll func(now, root int)
	dfsAll = func(now, root int) {
		for _, n := range tMap[now] {
			if n == root {
				continue
			}
			//子节点路径和=父节点路径和 - 子节点的子节点数量 + （N - 子节点的子节点数量 这个可以看作是右侧节点数量）
			ans[n] = ans[now] - childNums[n] + N - childNums[n]
			dfsAll(n, now)
		}
	}
	dfsAll(0, -1)
	return ans
}
