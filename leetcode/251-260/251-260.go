package _251_260

// leetcode problem No. 253
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
