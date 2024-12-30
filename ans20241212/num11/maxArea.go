package num11

func maxArea(height []int) (ans int) {
	i, j := 0, len(height)-1
	for i < j {
		ans = max(ans, min(height[i], height[j])*(j-i))
		if height[i] <= height[j] {
			i++
		} else {
			j--
		}
	}
	return
}
