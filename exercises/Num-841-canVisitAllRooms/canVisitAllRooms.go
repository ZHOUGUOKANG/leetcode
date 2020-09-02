package Num_841_canVisitAllRooms

func canVisitAllRooms(rooms [][]int) bool {
	visited := map[int]bool{}
	var queue []int
	queue = append(queue, rooms[0]...)
	visited[0] = true
	for len(queue) > 0 {
		pop := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		visited[pop] = true
		for _, v := range rooms[pop] {
			if _, ok := visited[v]; !ok {
				queue = append(queue, v)
			}
		}
	}
	return len(visited) == len(rooms)
}
