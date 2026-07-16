package _2671_2680

// leetcode No. 2673
func minIncrements(n int, cost []int) int {
	left, right := (n/2)/2, n/2
	ans := 0
	for right > left {
		for i := left; i < right; i++ {
			leftChild := i*2 + 1
			rightChild := i*2 + 2
			if cost[leftChild] > cost[rightChild] {
				ans += cost[leftChild] - cost[rightChild]
				cost[i] += cost[leftChild] // note: do not multiple 2
			} else {
				ans += cost[rightChild] - cost[leftChild]
				cost[i] += cost[rightChild] // note: do not multiple 2
			}
		}
		right = left
		left = left / 2
	}
	return ans
}
