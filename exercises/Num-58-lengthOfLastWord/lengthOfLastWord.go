package Num_58_LengthOfLastWord

func lengthOfLastWord(s string) int {
	n := len(s)
	isTail, hasWord := true, false
	tailID := n
	for i := n - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if isTail {
				tailID = i
			} else {
				return tailID - i - 1
			}
		} else {
			hasWord = true
			isTail = false
		}
	}
	if !hasWord {
		return 0
	}
	return tailID
}
