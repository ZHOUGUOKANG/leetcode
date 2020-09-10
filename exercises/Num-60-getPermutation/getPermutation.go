package Num_60_getPermutation

//func getPermutation(n int, k int) string {
//	cache := make(map[int]bool,n)
//	var ret [][]byte
//	tmp := make([]byte,0)
//	helper := func() {}
//	helper = func() {
//		if len(tmp) == n{
//			cp := make([]byte,n)
//			copy(cp,tmp)
//			ret = append(ret,cp)
//			if len(ret) == k{
//				return
//			}
//		}
//		for i:=1;i<=n;i++{
//			if has,ok := cache[i];ok && has{
//				continue
//			}
//			tmp = append(tmp,byte(i+48))
//			cache[i] =true
//			helper()
//			tmp = tmp[:len(tmp)-1]
//			cache[i] = false
//		}
//	}
//	helper()
//	return string(ret[k-1])
//}

func getPermutation(n int, k int) string {
	factorial := make(map[int]int, n+1)
	factorial[0] = 1
	for i := 1; i <= n; i++ {
		factorial[i] = i * factorial[i-1]
	}
	k--
	visited := make([]uint8, n+1)
	var ans []byte
	for i := 1; i <= n; i++ {
		num := k/factorial[n-i] + 1
		k = k % factorial[n-i]
		for z := 1; z <= n; z++ {
			if visited[z] == 0 {
				num--
				if num == 0 {
					ans = append(ans, byte(z+48))
					visited[z] = 1
					break
				}
			}
		}
	}
	return string(ans)
}
