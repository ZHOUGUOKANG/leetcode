package leetcodes

/**
10. 正则表达式匹配
给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。

'.' 匹配任意单个字符
'*' 匹配零个或多个前面的那一个元素
所谓匹配，是要涵盖 整个 字符串 s的，而不是部分字符串。

说明:

s 可能为空，且只包含从 a-z 的小写字母。
p 可能为空，且只包含从 a-z 的小写字母，以及字符 . 和 *。
示例 1:

输入:
s = "aa"
p = "a"
输出: false
解释: "a" 无法匹配 "aa" 整个字符串。
示例 2:

输入:
s = "aa"
p = "a*"
输出: true
解释: 因为 '*' 代表可以匹配零个或多个前面的那一个元素, 在这里前面的元素就是 'a'。因此，字符串 "aa" 可被视为 'a' 重复了一次。
示例 3:

输入:
s = "ab"
p = ".*"
输出: true
解释: ".*" 表示可匹配零个或多个（'*'）任意字符（'.'）。
示例 4:

输入:
s = "aab"
p = "c*a*b"
输出: true
解释: 因为 '*' 表示零个或多个，这里 'c' 为 0 个, 'a' 被重复一次。因此可以匹配字符串 "aab"。
示例 5:

输入:
s = "mississippi"
p = "mis*is*p*."
输出: false
*/
func IsMatch(s string, p string) bool {
	println(s, p)
	if s == p {
		return true
	}
	sLen := len(s)
	pLen := len(p)
	if sLen < 1 || pLen < 1 {
		return false
	}
	i, j := 0, 0
	flag := true
	for {
		if p[j] == 42 {
			println("匹配到*号")
			//星号前面一个如果是 . 代表所有单词，
			if p[j-1] == 46 {
				println("星号前面一个如果是 . 代表所有单词 需要判断是不是最后一位")
				if j < pLen-1 {
					//星号的位置不是最后一位 应该一直+i知道匹配后一个字母
					println("不是最后一位，一直+i直到相等，若没有相等的，则返回false")
					for i < sLen {
						if s[i] == p[j+1] {
							j++
							break
						} else {
							i++
						}
					}
					if i == sLen {
						flag = false
						break
					}
				} else {
					println("是最后一位，返回true")
					return true
				}
			} else {
				//如果星号不是最后一位， 应该找到星号的下一位置，然后看该位置和现在i之间数否满足星号条件
				if j < pLen-1 {

				}
				for i < sLen {
					if s[i] == p[j-1] {
						i++
					} else {
						j++
						break
					}
				}
				println("i++直到i不等于j-1", i, j)
				if i == sLen {
					if j >= pLen-1 {
						flag = true
					} else {
						flag = false
					}
					break
				}
			}
		} else if p[j] == 46 {
			// . 号 '.' 匹配任意单个字符
			i++
			j++
			println("匹配到.号，i++ j++", i, j)
		} else {
			//完全匹配
			if p[j] == s[i] {
				j++
				i++
				println("字母完全匹配，i++ j++", i, j)
			} else {
				if j+1 < pLen && p[j+1] == 42 {
					j++
					println("字母不完全匹配，但是字母后边接 * j++", j)
				} else {
					println("字母不完全匹配，返回false")
					flag = false
					break
				}
			}
		}
		if i >= sLen {
			if j >= pLen {
				flag = true
				break
			}
			if j < pLen-2 {
				flag = false
				break
			}
			if j == pLen-2 {
				j++
			}
			if p[j] == 42 || p[j] == 46 {
				flag = true
			} else {
				flag = false
			}
			break
		}
		if j >= pLen {
			if i >= sLen {
				flag = true
			} else {
				flag = false
			}
			break
		}
	}

	return flag
}
