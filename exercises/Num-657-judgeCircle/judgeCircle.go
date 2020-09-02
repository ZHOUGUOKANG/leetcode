package Num_657_judgeCircle

func judgeCircle(moves string) bool {
	i, j := 0, 0
	for _, v := range moves {
		switch v {
		case 'U':
			j++
		case 'D':
			j--
		case 'L':
			i--
		case 'R':
			i++
		}
	}
	return i == 0 && j == 0
}
