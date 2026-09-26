package _1801_1810

import "slices"

// leetcode problem No. 1802
func maxValue(n int, index int, maxSum int) int {
	sum := func(value int) int {
		ans := 0
		startIndex := max(0, index-value+1)
		if startIndex > 0 {
			ans += startIndex // add all the ones
		}
		numOfValue := index - startIndex + 1
		startValue := max(1, value-index)

		ans += (value + startValue) * numOfValue / 2

		endIndex := min(n-1, value+index-1)
		if endIndex < n-1 {
			ans += n - 1 - endIndex // add all the ones
		}
		numOfValue = endIndex - index + 1
		endValue := max(1, value-(n-1-index))
		ans += (value + endValue) * numOfValue / 2
		ans -= value
		return ans
	}

	right := maxSum
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

// leetcode problem No. 1807
func evaluate(s string, knowledge [][]string) string {
	knowledgeMap := make(map[string]string)
	for _, pair := range knowledge {
		knowledgeMap[pair[0]] = pair[1]
	}
	n := len(s)
	ans := make([]byte, 0, n)
	i := 0

	for i < n {
		if s[i] == '(' {
			j := i + 1
			for j < n && s[j] != ')' {
				j++
			}
			key := s[i+1 : j]
			if str, ok := knowledgeMap[key]; ok {
				ans = slices.Concat(ans, []byte(str))
			} else {
				ans = append(ans, '?')
			}
			i = j
		} else {
			ans = append(ans, s[i])
		}
		i++
	}

	return string(ans)
}
