package leetcodes

/**
14. 最长公共前缀
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

示例 1:

输入: ["flower","flow","flight"]
输出: "fl"
示例 2:

输入: ["dog","racecar","car"]
输出: ""
解释: 输入不存在公共前缀。
说明:

所有输入只包含小写字母 a-z 。

*/
func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	minLen := 0
	for key, value := range strs {
		if key == 0 {
			minLen = len(value)
			continue
		}
		if len(value) < minLen {
			minLen = len(value)
		}
	}
	tmpStr := ""
	flag := true
	for i := 0; i < minLen && flag; i++ {
		tmpPre := ""
		for key, value := range strs {
			if key == 0 {
				tmpPre = string(value[i])
				continue
			}
			if tmpPre != string(value[i]) {
				flag = false
				break
			}
			if key == len(strs)-1 && tmpPre == string(value[i]) {
				tmpStr += tmpPre
			}
		}
	}
	return tmpStr
}
