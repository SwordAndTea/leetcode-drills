package _1801_1810

// leetcode problem No. 1802

func maxValue(n int, index int, maxSum int) int {
	leftSum := (index + 1) * index / 2            // sum of left to pad
	rightSum := (n - 1 - index) * (n - index) / 2 // sum right to pad

	sum := func(value int) int {
		ans := 0
		startIndex := max(0, index-value+1)
		if startIndex > 0 {
			ans += startIndex
		}
		numOfValue := index - startIndex + 1
		startValue := max(1, value-index)

		ans += (value + startValue) * numOfValue / 2

		endIndex := min(n-1, value+index-1)
		if endIndex < n-1 {
			ans += n - 1 - endIndex
		}
		numOfValue = endIndex - index + 1
		endValue := max(1, value-(n-1-index))
		ans += (value + endValue) * numOfValue / 2
		ans -= value
		return ans
	}

	right := (leftSum + rightSum + maxSum) / n
	left := 1

	for left <= right {
		mid := (left + right) / 2
		midSum := sum(mid)

		if midSum <= maxSum && sum(mid+1) > maxSum {
			return mid
		}

		if midSum <= maxSum {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return left
}
