package Num_1239_MaxLength

import "math/bits"

func maxLength(arr []string) (ans int) {
	masks := []int{0} // 0 对应空串
outer:
	for _, s := range arr {
		mask := 0
		for _, ch := range s {
			ch -= 'a'
			if mask&(1<<ch) > 0 { // 若 mask 已有 ch，则说明 s 含有重复字母，无法构成可行解
				continue outer
			}
			mask |= 1 << ch // 将 ch 加入 mask 中
		}
		for _, m := range masks {
			if m&mask == 0 { // m 和 mask 无公共元素
				t := m | mask
				masks = append(masks, t)
				ans = max(ans, bits.OnesCount(uint(t)))
			}
		}
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
