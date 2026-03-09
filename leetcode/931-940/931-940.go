package _931_940

// leetcode problem No. 939

func minAreaRect(points [][]int) int {
	pointMap := map[int]map[int]bool{}

	for _, point := range points {
		if _, ok := pointMap[point[0]]; !ok {
			pointMap[point[0]] = map[int]bool{}
		}
		pointMap[point[0]][point[1]] = true
	}
	ans := -1
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			point1 := points[i]
			point2 := points[j]
			if point1[0] != point2[0] && point1[1] != point2[1] {
				minX := min(point1[0], point2[0])
				maxX := max(point1[0], point2[0])
				minY := min(point1[1], point2[1])
				maxY := max(point1[1], point2[1])
				if pointMap[minX][minY] && pointMap[maxX][maxY] && pointMap[minX][maxY] && pointMap[maxX][minY] {
					if area := (maxX - minX) * (maxY - minY); ans == -1 || area < ans {
						ans = area
					}
				}

			}
		}
	}
	if ans == -1 {
		return 0
	}
	return ans
}
