package Num_1047_removeDuplicates

func RemoveDuplicates(S string) string {
	//第一个推入栈，找下一个，下一个相同，推出栈顶，i++,不相同，推入栈顶,i++
	if len(S) <= 1 {
		return S
	}
	stack := []byte{S[0]}
	strLen := len(S)
	for i := 1; i < strLen; {
		if stack[len(stack)-1] == S[i] {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				i++
				if i >= strLen {
					break
				}
				stack = append(stack, S[i])
			}
		} else {
			stack = append(stack, S[i])
		}
		i++
	}
	return string(stack)
}
