package Num_88_Merge

//func merge(nums1 []int, m int, nums2 []int, n int)  {
//	copy(nums1[m:],nums2[:n])
//	sort.Ints(nums1)
//	return
//}

func merge(nums1 []int, m int, nums2 []int, n int) {
	l, r, p := m-1, n-1, m+n-1
	for l >= 0 && r >= 0 {
		if nums1[l] < nums2[r] {
			nums1[p] = nums2[r]
			r--
		} else {
			nums1[p] = nums1[l]
			l--
		}
		p--
	}
	copy(nums1[:r+1], nums2[:r+1])
}
