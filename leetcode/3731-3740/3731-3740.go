package _3731_3740

import "sort"

// leetcode problem No. 3731
func findMissingElements(nums []int) []int {
	sort.Ints(nums)
	n := len(nums)
	ans := make([]int, 0, nums[n-1]-nums[0])
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] > 1 {
			for j := nums[i-1] + 1; j < nums[i]; j++ {
				ans = append(ans, j)
			}
		}
	}
	return ans
}
