package Num_350_intersect

import "sort"

func intersect(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	offset1, offset2 := 0, 0
	result := make([]int, 0)
	for offset1 < len(nums1) && offset2 < len(nums2) {
		if nums1[offset1] == nums2[offset2] {
			result = append(result, nums1[offset1])
			offset1++
			offset2++
		} else if nums1[offset1] > nums2[offset2] {
			offset2++
		} else {
			offset1++
		}
	}
	return result
}
