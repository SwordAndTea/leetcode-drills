package _551_560

// leetcode problem No. 560

func subarraySum(nums []int, k int) int {
	// O(n^2) version
	//prefixSum := make([]int, len(nums))
	//sum := 0
	//for i := 0; i < len(nums); i++ {
	//	sum += nums[i]
	//	prefixSum[i] = sum
	//}
	//ans := 0
	//for i := 0; i < len(prefixSum); i++ {
	//	for j := i; j < len(prefixSum); j++ {
	//		if prefixSum[j]-prefixSum[]+nums[i] == k {
	//			ans++
	//		}
	//	}
	//}
	//return ans

	// O(n) version
	prefixSumCount := make(map[int]int) // prefixSumCount store how many times a prefix sum has occurred
	prefixSumCount[0] = 1
	sum := 0
	ans := 0
	for _, num := range nums {
		sum += num
		if count, ok := prefixSumCount[k-sum]; ok {
			ans += count
		}
		prefixSumCount[sum]++
	}
	return ans
}
