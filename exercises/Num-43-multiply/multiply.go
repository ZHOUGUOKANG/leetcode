package Num_43_multiply

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	if len(num1) < len(num2) {
		num1, num2 = num2, num1
	}
	rs1 := []byte{48}
	nv := []byte(num1)
	appendZero := make([]byte, 110)
	for j := 0; j < 110; j++ {
		appendZero[j] = 48
	}
	for i := len(num2) - 1; i >= 0; i-- {
		rs2 := helper(num2[i]-48, nv)
		if rs2 == nil {
			continue
		} else {
			tmpZeroNum := len(num2) - 1 - i
			rs2 = append(rs2, appendZero[:tmpZeroNum]...)
		}
		rs1 = bytesAdd(rs1, rs2)
	}
	return string(rs1)
}

func helper(m uint8, n []byte) []byte {
	if m == 0 {
		return nil
	} else if m == 1 {
		return n
	} else {
		var carry uint8 = 0
		nv := make([]byte, len(n)+1)
		for i := len(n) - 1; i >= 0; i-- {
			r := m*(n[i]-48) + carry
			carry = r / 10
			nv[i+1] = r%10 + 48
		}
		if carry == 0 {
			return nv[1:]
		} else {
			nv[0] = carry + 48
			return nv
		}
	}
}

func bytesAdd(a, b []byte) []byte {
	m, n := len(a), len(b)
	if m > n {
		m, n = n, m
		a, b = b, a
	}
	rs := make([]byte, n+1)
	var tmpA, carry, sum uint8 = 0, 0, 0
	for i, j := m-1, n-1; i >= 0 || j >= 0; i, j = i-1, j-1 {
		tmpA = carry
		if i >= 0 {
			tmpA += a[i] - 48
		}
		sum = tmpA + b[j] - 48
		carry = sum / 10
		rs[j+1] = sum%10 + 48
	}
	if carry == 0 {
		return rs[1:]
	} else {
		rs[0] = carry + 48
		return rs
	}
}
