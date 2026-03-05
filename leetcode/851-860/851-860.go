package _851_860

// leetcode problem No. 852

func peakIndexInMountainArray(arr []int) int {
	n := len(arr)
	left, right := 0, n-1
	for left <= right {
		mid := left + (right-left)/2
		if mid-1 >= 0 && mid+1 < n && arr[mid-1] < arr[mid] && arr[mid+1] < arr[mid] {
			return mid
		} else if (mid-1 >= 0 && arr[mid-1] < arr[mid]) || (mid+1 < n && arr[mid+1] > arr[mid]) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
