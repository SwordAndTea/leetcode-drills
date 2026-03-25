package _971_980

// leetcode problem No. 974

func subarraysDivByK(nums []int, k int) int {
	prefixModCount := make(map[int]int)
	curSum := 0
	prefixModCount[0] = 1
	ans := 0
	for _, num := range nums {
		curSum = (curSum + num) % k
		if curSum < 0 {
			curSum += k
		}
		ans += prefixModCount[curSum]
		prefixModCount[curSum]++
	}
	return ans
}
