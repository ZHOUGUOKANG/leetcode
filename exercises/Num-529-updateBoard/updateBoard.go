package Num_529_updateBoard

var m, n int

func updateBoard(board [][]byte, click []int) [][]byte {
	cM, cN := click[0], click[1]
	if board[cM][cN] == 'M' {
		board[cM][cN] = 'X'
		return board
	}
	m, n = len(board), len(board[0])
	helper(board, cM, cN)
	return board
}

func helper(board [][]byte, cM, cN int) {
	if cM >= m || cM < 0 || cN >= n || cN < 0 {
		return
	}
	target := board[cM][cN]
	if target != 'E' {
		return
	}
	board[cM][cN] = 'B'
	count := 0
	for i := cM - 1; i <= cM+1; i++ {
		for j := cN - 1; j <= cN+1; j++ {
			count += hasM(board, i, j)
		}
	}
	if count == 0 {
		for i := cM - 1; i <= cM+1; i++ {
			for j := cN - 1; j <= cN+1; j++ {
				helper(board, i, j)
			}
		}
	} else {
		board[cM][cN] = byte(count + 48)
	}
}

func hasM(board [][]byte, cM, cN int) int {
	if cM >= m || cM < 0 || cN >= n || cN < 0 {
		return 0
	}
	if board[cM][cN] == 'M' {
		return 1
	} else {
		return 0
	}
}
