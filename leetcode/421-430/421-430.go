package _421_430

// leetcode problem No. 424
func characterReplacement(s string, k int) int {
	freqMap := make(map[byte]int) // freq map record for s[left:right(include)]
	left := 0
	right := 0
	ans := 0
	maxFreq := 0
	for right < len(s) {
		freqMap[s[right]]++
		if freqMap[s[right]] > maxFreq {
			maxFreq = freqMap[s[right]]
		}
		for right-left+1 > maxFreq+k {
			freqMap[s[left]]--
			left++

			// update maxFreq
			// though we updated maxFreq, but we can skip this, the reason is:
			// k is fixed, only maxFreq become bigger, the final ans=right-left+1 will be updated
			maxFreq = 0
			for _, freq := range freqMap {
				if freq > maxFreq {
					maxFreq = freq
				}
			}
		}
		ans = max(ans, right-left+1)
		right++
	}
	return ans
}
