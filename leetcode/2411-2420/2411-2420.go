package _2411_2420

// leetcode problem No. 2420

func goodIndices(nums []int, k int) []int {
	n := len(nums)
	leftDp := make([]int, n)
	left := k - 1
	for left-1 >= 0 && nums[left-1] >= nums[left] {
		left--
	}
	leftDp[k-1] = k - 1 - left + 1

	for i := k; i < n-k; i++ {
		if nums[i] <= nums[i-1] {
			leftDp[i] = leftDp[i-1] + 1
		} else {
			leftDp[i] = 0
		}
	}

	rightDp := make([]int, n)
	right := n - k
	for right+1 < n && nums[right+1] >= nums[right] {
		right++
	}
	rightDp[n-k] = right - n + k + 1

	for i := n - k - 1; i >= k; i-- {
		if nums[i] <= nums[i+1] {
			rightDp[i] = rightDp[i+1] + 1
		} else {
			rightDp[i] = 0
		}
	}

	ans := make([]int, 0, k)
	for i := k; i < n-k; i++ {
		if leftDp[i-1] >= k && rightDp[i+1] >= k {
			ans = append(ans, i)
		}
	}

	return ans
}
