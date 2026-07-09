package _3531_3540

// leetcode problem 3532
func pathExistenceQueries(n int, nums []int, maxDiff int, queries [][]int) []bool {
	clusterInfo := make([]int, n)
	clusterID := 0
	clusterInfo[0] = 0

	for i := 1; i < n; i++ {
		if nums[i]-nums[i-1] > maxDiff {
			clusterID++
		}
		clusterInfo[i] = clusterID
	}

	m := len(queries)
	result := make([]bool, m)
	for i, q := range queries {
		if clusterInfo[q[0]] == clusterInfo[q[1]] {
			result[i] = true
		} else {
			result[i] = false
		}
	}

	return result
}
