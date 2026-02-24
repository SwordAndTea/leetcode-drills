package _251_260

//import "sort"

// leetcode problem No. 253

// brute force: O(n^2)
//func minMeetingRooms(intervals [][]int) int {
//	sort.Slice(intervals, func(i, j int) bool {
//		return intervals[i][1] < intervals[j][1]
//	})
//
//	ans := 1
//	for i := 0; i < len(intervals); i++ {
//		numRoomNeeded := 1
//		for j := i + 1; j < len(intervals); j++ {
//			if intervals[j][0] < intervals[i][1] {
//				numRoomNeeded++
//			} else {
//				break
//			}
//		}
//		if numRoomNeeded > ans {
//			ans = numRoomNeeded
//		}
//	}
//
//	return ans
//}

// difference array

func minMeetingRooms(intervals [][]int) int {
	m := 0
	for _, v := range intervals {
		m = max(m, v[1])
	}
	// d[i] stands for the number of meeting rooms we need to change at time i
	// then it will act like a difference array
	d := make([]int, m+1)
	for _, v := range intervals {
		d[v[0]]++ // increase a room at meeting start
		d[v[1]]-- // decrease a room at meeting end
	}
	prefixSum := 0
	ans := 0
	for _, v := range d {
		prefixSum += v
		if prefixSum > ans {
			ans = prefixSum
		}
	}
	return ans
}
