package Num_75_SortColors

func sortColors(nums []int) {
	l, r := 0, len(nums)-1
	for _, v := range []int{2, 1} {
		l = 0
		for l <= r {
			if nums[l] == v {
				nums[l], nums[r] = nums[r], nums[l]
				r--
			} else {
				l++
			}
		}
	}
}

/**

func sortColors(nums []int) {
    p0, p1 := 0, 0
    for i, c := range nums {
        if c == 0 {
            nums[i], nums[p0] = nums[p0], nums[i]
            if p0 < p1 {
                nums[i], nums[p1] = nums[p1], nums[i]
            }
            p0++
            p1++
        } else if c == 1 {
            nums[i], nums[p1] = nums[p1], nums[i]
            p1++
        }
    }
}
*/
