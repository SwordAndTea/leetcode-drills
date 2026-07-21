package _281_290

// leetcode problem No. 289
func gameOfLife(board [][]int) {
	m, n := len(board), len(board[0])
	nextState := make([][]int, m)
	for i := range nextState {
		nextState[i] = make([]int, n)
	}
	// do update
	verticalDirection := []int{-1, 0, 1}
	horizontalDirection := []int{-1, 0, 1}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			liveNeighbors := 0
			for _, vd := range verticalDirection {
				for _, hd := range horizontalDirection {
					if vd == 0 && hd == 0 {
						continue
					}
					neighborI := i + vd
					neighborJ := j + hd
					if neighborI >= 0 && neighborI < m && neighborJ >= 0 && neighborJ < n && board[neighborI][neighborJ] == 1 {
						liveNeighbors++
					}
				}
			}
			if board[i][j] == 1 {
				if liveNeighbors < 2 || liveNeighbors > 3 {
					nextState[i][j] = 0
				} else {
					nextState[i][j] = 1
				}
			} else {
				if liveNeighbors == 3 {
					nextState[i][j] = 1
				} else {
					nextState[i][j] = 0
				}
			}
		}
	}

	// copy back
	for i := 0; i < m; i++ {
		copy(board[i], nextState[i])
	}
}
