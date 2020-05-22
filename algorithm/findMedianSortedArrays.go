package leetcodes

/**
4. 寻找两个有序数组的中位数
给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。

请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。

你可以假设 nums1 和 nums2 不会同时为空。
*/
func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1, len2 := len(nums1), len(nums2)
	count := len1 + len2
	tmpArray := make([]int, 0)

	if len1 == 0 {
		tmpArray = nums2
	} else if len2 == 0 {
		tmpArray = nums1
	} else {
		num1, num2 := 0, 0
		for num1 < len1 || num2 < len2 {
			if nums1[num1] < nums2[num2] {
				tmpArray = append(tmpArray, nums1[num1])
				num1++
				if num1 == len1 {
					tmpArray = append(tmpArray, nums2[num2:]...)
					break
				}
			} else {
				tmpArray = append(tmpArray, nums2[num2])
				num2++
				if num2 == len2 {
					tmpArray = append(tmpArray, nums1[num1:]...)
					break
				}
			}
		}
	}
	if count%2 == 0 {
		return float64(tmpArray[count/2-1]+tmpArray[count/2]) / 2
	} else {
		return float64(tmpArray[count/2])
	}
}
