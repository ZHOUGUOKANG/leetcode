package Num_204_countPrimes

func countPrimes(n int) (count int) {
	flag := make([]uint64, n/64+1)
	for i := range flag {
		flag[i] = 0
	}
	for i := 2; i < n; i++ {
		if (flag[celi(i)] & (1 << (i % 64))) > 0 {
			continue
		}
		count++
		for j := i * 2; j < n; j += i {
			flag[celi(j)] = flag[celi(j)] | (1 << j % 64)
		}
	}
	return
}

func celi(n int) int {
	if n < 64 {
		return 0
	}
	if n%64 > 0 {
		return n/64 + 1
	}
	return n / 64
}
