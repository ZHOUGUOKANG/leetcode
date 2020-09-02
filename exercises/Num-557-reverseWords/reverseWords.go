package Num_557_reverseWords

func reverseWords(s string) string {
	t := []rune(s)
	reverse := func(start, end int) {
		for start < end {
			t[start], t[end] = t[end], t[start]
			start++
			end--
		}
	}
	l, r := 0, 0
	for r < len(t) {
		if t[r] == ' ' {
			reverse(l, r-1)
			l = r + 1
		}
		r++
	}
	reverse(l, len(t)-1)
	return string(t)
}
