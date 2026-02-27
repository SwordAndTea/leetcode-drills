package _1451_1460

// leetcode problem No. 1458
func maxDotProduct(nums1 []int, nums2 []int) int {
	// dp[i][j] stands for the max dot product of the subsequences nums1[0:i+1] and nums2[0:j+1]
	dp := make([][]int, len(nums1))
	for i := 0; i < len(nums1); i++ {
		dp[i] = make([]int, len(nums2))
	}

	dp[0][0] = nums1[0] * nums2[0]

	for i := 1; i < len(nums1); i++ {
		dp[i][0] = max(nums1[i]*nums2[0], dp[i-1][0])
	}

	for j := 1; j < len(nums2); j++ {
		dp[0][j] = max(nums1[0]*nums2[j], dp[0][j-1])
	}

	for i := 1; i < len(nums1); i++ {
		for j := 1; j < len(nums2); j++ {
			// dp[i-1][j] already include situation that nums[i-1] is multiplying with nums[j]
			// and it's the same for dp[i][j-1]
			dp[i][j] = max(
				nums1[i]*nums2[j],
				dp[i-1][j-1],
				dp[i-1][j-1]+nums1[i]*nums2[j],
				dp[i-1][j],
				dp[i][j-1],
			)
		}
	}

	return dp[len(nums1)-1][len(nums2)-1]
}
