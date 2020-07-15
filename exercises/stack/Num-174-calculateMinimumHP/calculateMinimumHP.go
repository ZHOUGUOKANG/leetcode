package Num_174_calculateMinimumHP

func calculateMinimumHP(dungeon [][]int) int {
	hangLen := len(dungeon)
	lieLen := len(dungeon[0])
	loss := 1
	if dungeon[hangLen-1][lieLen-1] < 0 {
		dungeon[hangLen-1][lieLen-1] = 1 - dungeon[hangLen-1][lieLen-1]
	} else {
		dungeon[hangLen-1][lieLen-1] = 1
	}
	for i := hangLen - 1; i >= 0; i-- {
		for j := lieLen - 1; j >= 0; j-- {
			if i == hangLen-1 && j == lieLen-1 {
				continue
			}
			if i == hangLen-1 {
				loss = dungeon[i][j+1]
			} else if j == lieLen-1 {
				loss = dungeon[i+1][j]
			} else {
				if dungeon[i+1][j] > dungeon[i][j+1] {
					loss = dungeon[i][j+1]
				} else {
					loss = dungeon[i+1][j]
				}
			}

			if dungeon[i][j] == 0 {
				dungeon[i][j] = loss
			} else if dungeon[i][j] > 0 {
				if dungeon[i][j] >= loss {
					dungeon[i][j] = 1
				} else {
					dungeon[i][j] = loss - dungeon[i][j]
				}
			} else {
				dungeon[i][j] = loss - dungeon[i][j]
			}
		}
	}

	return dungeon[0][0]
}
