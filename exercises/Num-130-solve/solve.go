package Num_130_solve

func solve(board [][]byte) {
	if len(board) < 3 || len(board[0]) < 3 {
		return
	}
	stack := make([][]int, 0)
	m, n := len(board), len(board[0])
	for i := 0; i < m; i++ {
		if i == 0 || i == m-1 {
			for j := 0; j < n; j++ {
				if board[i][j] == 'O' {
					stack = append(stack, []int{i, j})
				}
			}
		} else {
			if board[i][0] == 'O' {
				stack = append(stack, []int{i, 0})
			}
			if board[i][n-1] == 'O' {
				stack = append(stack, []int{i, n - 1})
			}
		}
	}
	addOne := func(i, j int) {
		if i > -1 && i < m && j > -1 && j < n && board[i][j] == 'O' {
			stack = append(stack, []int{i, j})
		}
	}
	for len(stack) > 0 {
		pop := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		i, j := pop[0], pop[1]
		board[i][j] = 'Y'
		addOne(i-1, j)
		addOne(i+1, j)
		addOne(i, j-1)
		addOne(i, j+1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == 'Y' {
				board[i][j] = 'O'
			}
		}
	}
}
