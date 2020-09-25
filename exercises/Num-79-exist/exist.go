package Num_79_exist

func exist(board [][]byte, word string) bool {
	m := len(board)
	if m == 0 {
		return false
	}
	n := len(board[0])
	wl := len(word)
	helper := func(i, j, w int) bool { return false }
	helper = func(i, j, w int) bool {
		if w == wl {
			return true
		}
		if i < 0 || i >= m || j < 0 || j >= n {
			return false
		}
		if board[i][j] != word[w] || board[i][j] == 'z' {
			return false
		}
		t := board[i][j]
		board[i][j] = 'z'
		w++
		if helper(i+1, j, w) || helper(i-1, j, w) || helper(i, j-1, w) || helper(i, j+1, w) {
			return true
		}
		board[i][j] = t
		return false
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[0] && helper(i, j, 0) {
				return true
			}
		}
	}
	return false
}
