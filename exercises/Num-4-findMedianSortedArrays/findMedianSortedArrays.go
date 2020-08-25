package Num_4_findMedianSortedArrays

import (
	"math"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	sum := len(nums1) + len(nums2)
	if sum%2 == 1 {
		return find(nums1[:], nums2[:], sum/2+1)
	} else {
		return (find(nums1[:], nums2[:], sum/2) + find(nums1[:], nums2[:], sum/2+1)) / 2
	}
}

func find(nums1, nums2 []int, k int) float64 {
	if len(nums1) == 0 {
		return float64(nums2[k-1])
	}
	if len(nums2) == 0 {
		return float64(nums1[k-1])
	}
	if k == 1 {
		return math.Min(float64(nums1[0]), float64(nums2[0]))
	}
	offset := k / 2
	offset1, offset2 := offset-1, offset-1
	if len(nums1) < offset {
		offset1 = len(nums1) - 1
	}
	if len(nums2) < offset {
		offset2 = len(nums2) - 1
	}
	if nums1[offset1] > nums2[offset2] {
		return find(nums1, nums2[offset2+1:], k-offset2-1)
	} else {
		return find(nums1[offset1+1:], nums2, k-offset1-1)
	}
}
