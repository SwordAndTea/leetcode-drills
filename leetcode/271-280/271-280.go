package _271_280

import "sort"

// leetcode problem No. 274

func hIndex(citations []int) int {
	sort.Ints(citations)
	ans := 0
	n := len(citations)
	for i := 0; i < n; i++ {
		if c := min(citations[i], n-i); c > ans {
			ans = c
		}
	}
	return ans
}
