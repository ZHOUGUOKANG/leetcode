package Num_835_LargestOverlap

func largestOverlap(img1 [][]int, img2 [][]int) int {
	n, maxNum := len(img1), 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			t1, t2, t3, t4 := 0, 0, 0, 0
			for x := 0; x < n; x++ {
				for y := 0; y < n; y++ {
					if img1[x][y] == 1 {
						if x+i < n && y+j < n && img2[x+i][y+j] == 1 {
							t1++
						}
						if x+i < n && y-j > -1 && img2[x+i][y-j] == 1 {
							t2++
						}
						if x-i > -1 && y+j < n && img2[x-i][y+j] == 1 {
							t3++
						}
						if x-i > -1 && y-j > -1 && img2[x-i][y-j] == 1 {
							t4++
						}
					}
				}
			}
			maxNum = max(maxNum, max(max(t1, t2), max(t3, t4)))
		}
	}
	return maxNum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
