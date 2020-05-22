package sort

/**
选择排序

选择排序(Selection-sort)是一种简单直观的排序算法。
它的工作原理：首先在未排序序列中找到最小（大）元素，存放到排序序列的起始位置，然后，再从剩余未排序元素中继续寻找最小（大）元素，然后放到已排序序列的末尾。
以此类推，直到所有元素均排序完毕。
*/
func SelectionSort(s []int) []int {
	tmpIndex := 0
	tmp := 0
	maxIndex := 0
	length := len(s)
	for i := 0; i < length-1; i++ {
		maxIndex = length - 1 - i
		tmpIndex = 0
		for j := 0; j < maxIndex; j++ {
			if s[tmpIndex] < s[j] {
				tmpIndex = j
			}
		}
		if s[tmpIndex] > s[maxIndex] {
			tmp = s[tmpIndex]
			s[tmpIndex] = s[maxIndex]
			s[maxIndex] = tmp
		}
	}
	return s
}
