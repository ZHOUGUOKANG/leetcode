package Num_38_countAndSay

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	t1 := []byte{'1'}
	for i := 1; i < n; i++ {
		var t []byte
		l, r := 0, 0
		for l < len(t1) && r < len(t1) {
			if t1[l] == t1[r] {
				r++
			} else {
				t = append(t, byte(r-l+48), t1[l])
				l = r
			}
		}
		if l != r {
			t = append(t, byte(r-l+48), t1[l])
		}
		t1 = t
	}
	return string(t1)
}
