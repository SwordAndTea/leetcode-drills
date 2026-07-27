package _2531_2540

// leetcode problem No. 2536
func rangeAddQueries(n int, queries [][]int) [][]int {
	grid := make([][]int, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]int, n)
	}
	for _, query := range queries {
		for i := query[0]; i <= query[2]; i++ { //for each row
			grid[i][query[1]]++
			if query[3]+1 < n {
				grid[i][query[3]+1]--
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 1; j < n; j++ {
			grid[i][j] += grid[i][j-1]
		}
	}
	return grid
}
