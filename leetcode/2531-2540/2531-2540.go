package _2531_2540

// leetcode problem No. 2534
// https://github.com/doocs/leetcode/blob/main/solution/2500-2599/2534.Time%20Taken%20to%20Cross%20the%20Door/README_EN.md
func timeTaken(arrival []int, state []int) []int {
	n := len(arrival)
	ans := make([]int, n)
	indexWaitForEnter := make([]int, 0)
	indexWaitForExit := make([]int, 0)
	curT := arrival[0]
	scanIndex := 0
	prevState := -1
	for scanIndex < n || len(indexWaitForEnter) > 0 || len(indexWaitForExit) > 0 {
		for scanIndex < n && arrival[scanIndex] <= curT {
			if state[scanIndex] == 0 {
				indexWaitForEnter = append(indexWaitForEnter, scanIndex)
				scanIndex++
			} else {
				indexWaitForExit = append(indexWaitForExit, scanIndex)
				scanIndex++
			}
		}

		if prevState == -1 || prevState == 1 {
			if len(indexWaitForExit) > 0 {
				index := indexWaitForExit[0]
				indexWaitForExit = indexWaitForExit[1:]
				ans[index] = curT
				prevState = 1
			} else if len(indexWaitForEnter) > 0 {
				index := indexWaitForEnter[0]
				indexWaitForEnter = indexWaitForEnter[1:]
				ans[index] = curT
				prevState = 0
			}
		} else { // if prevState == 0
			if len(indexWaitForEnter) > 0 {
				index := indexWaitForEnter[0]
				indexWaitForEnter = indexWaitForEnter[1:]
				ans[index] = curT
				prevState = 0
			} else if len(indexWaitForExit) > 0 {
				index := indexWaitForExit[0]
				indexWaitForExit = indexWaitForExit[1:]
				ans[index] = curT
				prevState = 1
			}
		}
		curT++
	}

	return ans
}

// leetcode problem No. 2536
func rangeAddQueries(n int, queries [][]int) [][]int {
	grid := make([][]int, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]int, n)
	}
	for _, query := range queries {
		for i := query[0]; i <= query[2]; i++ { //for each row
			grid[i][query[1]]++
			if end := query[3] + 1; end < n { // note: add 1 here
				grid[i][end]--
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
