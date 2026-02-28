package _991_1000

import "math"

// leetcode problem No. 994

func orangesRotting(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	rottenTime := make([][]int, m)
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		rottenTime[i] = make([]int, n)
		for j := 0; j < n; j++ {
			rottenTime[i][j] = math.MaxInt
		}
		visited[i] = make([]bool, n)
	}

	bfs := func(startI, startJ int) {
		q := make([][2]int, 1)
		q[0] = [2]int{startI, startJ}
		visited[startI][startJ] = true
		minutes := 0
		for len(q) > 0 {
			minutes++
			for i := len(q); i > 0; i-- {
				curI := q[0][0]
				curJ := q[0][1]
				q = q[1:]

				if curI-1 >= 0 && grid[curI-1][curJ] == 1 && !visited[curI-1][curJ] {
					q = append(q, [2]int{curI - 1, curJ})
					rottenTime[curI-1][curJ] = min(minutes, rottenTime[curI-1][curJ])
					visited[curI-1][curJ] = true
				}

				if curI+1 < m && grid[curI+1][curJ] == 1 && !visited[curI+1][curJ] {
					q = append(q, [2]int{curI + 1, curJ})
					rottenTime[curI+1][curJ] = min(minutes, rottenTime[curI+1][curJ])
					visited[curI+1][curJ] = true
				}

				if curJ-1 >= 0 && grid[curI][curJ-1] == 1 && !visited[curI][curJ-1] {
					q = append(q, [2]int{curI, curJ - 1})
					rottenTime[curI][curJ-1] = min(minutes, rottenTime[curI][curJ-1])
					visited[curI][curJ-1] = true
				}

				if curJ+1 < n && grid[curI][curJ+1] == 1 && !visited[curI][curJ+1] {
					q = append(q, [2]int{curI, curJ + 1})
					rottenTime[curI][curJ+1] = min(minutes, rottenTime[curI][curJ+1])
					visited[curI][curJ+1] = true
				}
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				bfs(i, j)
				for k := 0; k < m; k++ {
					clear(visited[k])
				}
			}
		}
	}
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				if rottenTime[i][j] == math.MaxInt {
					return -1
				}
				ans = max(ans, rottenTime[i][j])
			}
		}
	}

	return ans
}
