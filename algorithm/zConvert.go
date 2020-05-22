package leetcodes

/**
6. Z 字形变换
将一个给定字符串根据给定的行数，以从上往下、从左到右进行 Z 字形排列。

比如输入字符串为 "LEETCODEISHIRING" 行数为 3 时，排列如下：

L   C   I   R
E T O E S I I G
E   D   H   N
之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："LCIRETOESIIGEDHN"。
*/
func ZConvert(s string, numRows int) string {
	if s == "" || numRows <= 0 {
		return ""
	}
	if numRows == 1 {
		return s
	}
	matrix := make([][]byte, numRows)
	row := 0
	fangxiang := -1
	for i := 0; i < len(s); i++ {
		matrix[row] = append(matrix[row], s[i])
		if fangxiang == -1 {
			row++
		} else {
			row--
		}
		if row >= numRows || row < 0 {
			fangxiang = -fangxiang
			if fangxiang == -1 {
				row += 2
			} else {
				row -= 2
			}
		}
	}
	res := ""
	for i := 0; i < numRows; i++ {
		res += string(matrix[i])
	}
	return res
}
