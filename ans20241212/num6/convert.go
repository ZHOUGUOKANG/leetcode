package num6

import "bytes"

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	bs := make([][]byte, numRows)
	n, k := 0, -1
	for i := 0; i < len(s); i++ {
		bs[n] = append(bs[n], s[i])
		if n == 0 || n == numRows-1 {
			k = -k
		}
		n += k
	}
	return string(bytes.Join(bs, nil))
}
