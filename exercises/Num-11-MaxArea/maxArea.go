package Num_11_MaxArea

func maxArea(height []int) int {
	area, max := 0, 0
	for i, j := 0, len(height)-1; i < j; {
		if height[i] > height[j] {
			area = height[j] * (j - i)
			j--
		} else {
			area = height[i] * (j - i)
			i++
		}
		if area > max {
			max = area
		}
	}
	return max
}
