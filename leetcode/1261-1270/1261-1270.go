package _1261_1270

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func minTimeTwoPoints(point1, point2 []int) int {
	xAbs := abs(point1[0] - point2[0])
	yAbs := abs(point1[1] - point2[1])
	return max(xAbs, yAbs)
}

// leetcode problem No. 1266
func minTimeToVisitAllPoints(points [][]int) int {
	if len(points) <= 1 {
		return 0
	}
	time := 0
	for i := 1; i < len(points); i++ {
		time += minTimeTwoPoints(points[i-1], points[i])
	}
	return time
}
