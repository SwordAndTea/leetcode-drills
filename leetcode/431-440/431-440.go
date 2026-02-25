package _431_440

import "sort"

// leetcode problem No. 435

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	curInterval := intervals[0]
	remain := 1
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= curInterval[1] {
			remain++
			curInterval = intervals[i]
		}
	}
	return len(intervals) - remain
}
