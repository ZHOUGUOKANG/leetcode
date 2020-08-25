package Num_1209_removeDuplicates

func removeDuplicates(s string, k int) string {
	left, right := 0, 1
	stack := make([]uint8, 0)
	flag := true
	for right < len(s) {
		if s[left] == s[right] {
			if right-left+1 == k {
				right++
				left = right
				flag = false
			}
		} else {
			stack = append(stack, s[left:right]...)
			left = right
		}
		right++
	}
	if right == len(s) {
		stack = append(stack, s[left:right]...)
	}
	if flag {
		return string(stack)
	} else {
		return removeDuplicates(string(stack), k)
	}
}
