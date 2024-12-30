package num20

func isValid(s string) bool {
	stack := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(', '[', '{':
			stack = append(stack, s[i])
		case ')', ']', '}':
			if len(stack) == 0 {
				return false
			}
			if !isMatch(stack[len(stack)-1], s[i]) {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func isMatch(a, b byte) bool {
	return a == '(' && b == ')' || a == '[' && b == ']' || a == '{' && b == '}'
}
