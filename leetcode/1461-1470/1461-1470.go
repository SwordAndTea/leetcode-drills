package _1461_1470

import "math"

// leetcode problem No. 1464
func maxProduct(nums []int) int {
	ans := math.MinInt32
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if v := (nums[i] - 1) * (nums[j] - 1); v > ans {
				ans = v
			}
		}
	}
	return ans
}
