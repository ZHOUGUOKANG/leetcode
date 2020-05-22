package sort

/**
插入排序
插入排序（Insertion-Sort）的算法描述是一种简单直观的排序算法。
它的工作原理是通过构建有序序列，对于未排序数据，在已排序序列中从后向前扫描，找到相应位置并插入。
*/
func InsertionSort(s []int) []int {
	length := len(s)
	if length <= 1 {
		return s
	}
	tmpValue := 0
	for i := 1; i < length; i++ {
		if s[i] > s[i-1] {
			continue
		} else {
			tmpValue = s[i]
			j := i - 1
			for ; j >= 0 && s[j] > tmpValue; j-- {
				s[j+1] = s[j]
			}
			s[j+1] = tmpValue
		}
	}
	return s
}
