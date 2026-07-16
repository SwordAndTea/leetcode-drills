package _461_470

// leetcode problem No. 463
func islandPerimeter(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	perimeter := 0
	bfs := func(x, y int) {
		queue := [][2]int{{x, y}}
		visited[x][y] = true

		tryVisit := func(xx, yy int) {
			if xx < 0 || xx >= m || yy < 0 || yy >= n {
				perimeter++
				return
			}
			if grid[xx][yy] == 0 {
				perimeter++
				return
			}
			if !visited[xx][yy] {
				queue = append(queue, [2]int{xx, yy})
				visited[xx][yy] = true
			}
		}

		for len(queue) > 0 {
			cur := queue[0]
			queue = queue[1:]

			curX, curY := cur[0], cur[1]

			tryVisit(curX+1, curY)
			tryVisit(curX-1, curY)
			tryVisit(curX, curY+1)
			tryVisit(curX, curY-1)
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				bfs(i, j)
				return perimeter
			}
		}
	}

	return perimeter
}
