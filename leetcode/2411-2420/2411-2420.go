package _2411_2420

// leetcode problem No. 2420
func goodIndices(nums []int, k int) []int {
	n := len(nums)
	dp1 := make([]int, n) // dp1[i] stands for the length of non-increasing subarray ends with nums[i]
	dp2 := make([]int, n) // dp2[i] stands for the length of non-decreasing subarray starts with nums[i]

	dp1[0] = 1
	for i := 1; i < n; i++ {
		if nums[i] <= nums[i-1] {
			dp1[i] = dp1[i-1] + 1
		} else {
			dp1[i] = 1
		}
	}

	dp2[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		if nums[i] <= nums[i+1] {
			dp2[i] = dp2[i+1] + 1
		} else {
			dp2[i] = 1
		}
	}
	ans := make([]int, 0, k)
	for i := k; i < n-k; i++ {
		if dp1[i-1] >= k && dp2[i+1] >= k {
			ans = append(ans, i)
		}
	}
	return ans
}
