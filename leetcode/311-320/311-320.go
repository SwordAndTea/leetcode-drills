package _311_320

import "math"

// leetcode problem No. 316

func removeDuplicateLetters(s string) string {
	var stack []rune
	inStack := map[rune]bool{}
	lastIndex := map[rune]int{}

	for i, c := range s {
		lastIndex[c] = i
	}

	for i, c := range s {
		if !inStack[c] {
			// if the top char in the stack is greater than current char and it will appear after
			// pop it from the stack
			for len(stack) > 0 && c < stack[len(stack)-1] && i < lastIndex[stack[len(stack)-1]] {
				inStack[stack[len(stack)-1]] = false
				stack = stack[:len(stack)-1] // pop
			}
			inStack[c] = true
			stack = append(stack, c) // push
		}
	}
	return string(stack)
}

// leetcode problem No. 317

func shortestDistance(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	totalBuilding := 0
	totalDistToBuildings := make([][]int, m)   // the total distance for an empty land to all buildings
	reachableBuildingCount := make([][]int, m) // how many buildings can an empty land reach
	visited := make([][]bool, m)

	for i := 0; i < m; i++ {
		totalDistToBuildings[i] = make([]int, n)
		reachableBuildingCount[i] = make([]int, n)
		visited[i] = make([]bool, n)
	}

	var bfs func(int, int, [][]bool)
	bfs = func(startI, startJ int, visitMap [][]bool) {
		q := [][2]int{{startI, startJ}}
		visitMap[startI][startJ] = true
		distance := 0
		for len(q) > 0 {
			distance++
			// bfs travel by level
			for k := len(q); k > 0; k-- {
				curI := q[0][0]
				curJ := q[0][1]
				q = q[1:]
				if curI-1 >= 0 && grid[curI-1][curJ] == 0 && !visitMap[curI-1][curJ] {
					q = append(q, [2]int{curI - 1, curJ})
					visitMap[curI-1][curJ] = true
					totalDistToBuildings[curI-1][curJ] += distance
					reachableBuildingCount[curI-1][curJ] += 1
				}

				if curI+1 < m && grid[curI+1][curJ] == 0 && !visitMap[curI+1][curJ] {
					q = append(q, [2]int{curI + 1, curJ})
					visitMap[curI+1][curJ] = true
					totalDistToBuildings[curI+1][curJ] += distance
					reachableBuildingCount[curI+1][curJ] += 1
				}

				if curJ-1 >= 0 && grid[curI][curJ-1] == 0 && !visitMap[curI][curJ-1] {
					q = append(q, [2]int{curI, curJ - 1})
					visitMap[curI][curJ-1] = true
					totalDistToBuildings[curI][curJ-1] += distance
					reachableBuildingCount[curI][curJ-1] += 1
				}

				if curJ+1 < n && grid[curI][curJ+1] == 0 && !visitMap[curI][curJ+1] {
					q = append(q, [2]int{curI, curJ + 1})
					visitMap[curI][curJ+1] = true
					totalDistToBuildings[curI][curJ+1] += distance
					reachableBuildingCount[curI][curJ+1] += 1
				}
			}

		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				bfs(i, j, visited)
				for _, v := range visited {
					clear(v)
				}
				totalBuilding++
			}
		}
	}

	minDist := math.MaxInt
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if totalDistToBuildings[i][j] != 0 && reachableBuildingCount[i][j] == totalBuilding {
				minDist = min(minDist, totalDistToBuildings[i][j])
			}
		}
	}

	if minDist == math.MaxInt {
		return -1
	}
	return minDist
}
