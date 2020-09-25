package Num_offer_17_13_respace

//构建倒序字典树，
/**

假设当前我们已经考虑完了前 i 个字符了，对于前 i + 1 个字符对应的最少未匹配数：

第 i + 1 个字符未匹配，则 dp[i + 1] = dp[i] + 1，即不匹配数加 1;
遍历前 i 个字符，若以其中某一个下标 idx 为开头、以第 i + 1 个字符为结尾的字符串正好在词典里，则 dp[i] = min(dp[i], dp[idx]) 更新 dp[i]。

对于上述解法，计算 dp[i + 1]时，我们需要用 idx 来遍历前 i 个字符，逐个判断以 idx 为开头，以第 i + 1 个字符为结尾的字符串是否在字典里。
这一步可以利用字典树来加速，通过字典树我们可以查询以第 i + 1 个字符为结尾的单词有哪些（构建字典树时将单词逆序插入即可）。

*/

func respace(dictionary []string, sentence string) int {
	length := len(sentence)
	dictionaryMap := make(map[string]bool, len(dictionary))
	maxLength := 0
	for _, value := range dictionary {
		dictionaryMap[value] = true
		if len(value) > maxLength {
			maxLength = len(value)
		}
	}
	result := make([]int, length+1)
	for i := 1; i <= length; i++ {
		result[i] = result[i-1] + 1
		j := min(i-maxLength, 0)
		if j < 0 {
			j = 0
		}
		for ; j < i; j++ {
			if _, ok := dictionaryMap[sentence[j:i]]; ok {
				result[i] = min(result[i], result[j])
			}
		}
	}
	return result[length]
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
