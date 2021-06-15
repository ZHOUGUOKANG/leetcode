package Num_852_PeakIndexInMountainArray

func peakIndexInMountainArray(arr []int) int {
	maxIndex, maxValue := 0, 0
	for index, val := range arr {
		if val > maxValue {
			maxIndex, maxValue = index, val
		}
	}
	return maxIndex
}
func peakIndexInMountainArray2(arr []int) int {
	for i := 1; i < len(arr)-1; i++ {
		if arr[i] > arr[i-1] && arr[i] > arr[i+1] {
			return i
		}
	}
	return 1
}

func peakIndexInMountainArray1(arr []int) int {
	if len(arr) == 3 {
		return 1
	}
	n := len(arr)
	left, right := 0, n-1
	for left < right {
		mid := (left + right) / 2
		if right-left == 1 {
			if arr[right] > arr[left] {
				return right
			}
			return left
		} else {
			if mid-left >= 1 {
				if arr[mid] >= arr[mid-1] {
					left = mid
				} else {
					right = mid
				}
			} else {
				if arr[mid] >= arr[mid+1] {
					right = mid
				} else {
					left = mid
				}
			}
		}
	}
	return 0
}
