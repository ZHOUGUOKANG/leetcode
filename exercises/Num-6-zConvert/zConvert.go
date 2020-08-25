package Num_6_zConvert

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
