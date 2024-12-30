package num32

func longestValidParentheses(s string) (ans int) {
	stack := []int{-1}
	for i, x := range s {
		if x == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				ans = max(ans, i-stack[len(stack)-1])
			}
		}
	}
	return
}
