package Num_733_floodFill

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	targetColor := image[sr][sc]
	if targetColor == newColor {
		return image
	}
	h, l := len(image), len(image[0])
	helper := func(sr, sc int) {}
	helper = func(sr, sc int) {
		if sr >= 0 && sr < h && sc >= 0 && sc < l {
			if image[sr][sc] == targetColor {
				image[sr][sc] = newColor
				helper(sr-1, sc)
				helper(sr+1, sc)
				helper(sr, sc-1)
				helper(sr, sc+1)
			}
		}
	}
	helper(sr, sc)
	return image
}
