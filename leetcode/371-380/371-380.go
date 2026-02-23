package _71_380

// leetcode problem No. 347
func topKFrequent(nums []int, k int) []int {
	numCounter := make(map[int]int)
	maxFreq := 0
	for _, num := range nums {
		numCounter[num]++
		if numCounter[num] > maxFreq {
			maxFreq = numCounter[num]
		}
	}

	buckets := make([][]int, maxFreq+1)
	for num, count := range numCounter {
		buckets[count] = append(buckets[count], num)
	}

	result := make([]int, 0, k)
	for i := maxFreq; i >= 1; i-- {
		bucket := buckets[i]
		if len(bucket) != 0 {
			if len(result)+len(bucket) < k {
				result = append(result, bucket...)
			} else {
				result = append(result, bucket[:k-len(result)]...)
			}
		}
	}
	return result
}
