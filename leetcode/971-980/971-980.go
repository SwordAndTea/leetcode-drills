package _971_980

// leetcode problem No. 974
func subarraysDivByK(nums []int, k int) int {
	prefixModCount := make(map[int]int)
	curSumMod := 0
	prefixModCount[0] = 1
	ans := 0
	for _, num := range nums {
		curSumMod = (curSumMod + num) % k
		if curSumMod < 0 {
			curSumMod += k
		}
		ans += prefixModCount[curSumMod] //note: not k-curSumMod, as we need the substraction value be 0
		prefixModCount[curSumMod]++
	}
	return ans
}
