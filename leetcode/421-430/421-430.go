package _421_430

// leetcode problem No. 424

func characterReplacement(s string, k int) int {
	freqMap := map[byte]int{}
	i := 0
	j := 0
	ans := 0
	maxFreq := 0
	for j < len(s) {
		freqMap[s[j]]++
		if freqMap[s[j]] > maxFreq {
			maxFreq = freqMap[s[j]]
		}
		for j-i+1-maxFreq > k {
			// the reason why we don't update the maxFreq is that the possible max value is maxFreq + k
			// if there is no new maxFreq, the ans will not updated
			freqMap[s[i]]--
			i++
		}
		ans = max(ans, j-i+1)
		j++
	}
	return ans
}
