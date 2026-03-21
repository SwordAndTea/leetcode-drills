package _3001_3010

// leetcode problem No. 3004

func maxFrequencyElements(nums []int) int {
	maxFreq := 0
	freqMap := make(map[int]int)
	ans := 0
	for _, n := range nums {
		freqMap[n]++
		if freqMap[n] > maxFreq {
			maxFreq = freqMap[n]
			ans = maxFreq
		} else if freqMap[n] == maxFreq {
			ans += maxFreq
		}
	}

	return ans
}
