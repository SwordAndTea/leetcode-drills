package _51_60

func solveNQueens(n int) [][]string {
	result := make([][]string, 0)

	generateResultFromPermutation := func(permutation []int) []string {
		board := make([]string, 0, n)
		for i := 0; i < n; i++ {
			oneResult := make([]byte, n)
			for j := 0; j < n; j++ {
				if j == permutation[i] {
					oneResult[j] = 'Q'
				} else {
					oneResult[j] = '.'
				}
			}
			board = append(board, string(oneResult))
		}
		return board
	}

	abs := func(v int) int {
		if v < 0 {
			return -v
		}
		return v
	}

	isValidBoard := func(permutation []int) bool {
		for i := 0; i < len(permutation); i++ {
			for j := i + 1; j < len(permutation); j++ {
				if j-i == abs(permutation[i]-permutation[j]) {
					return false
				}
			}
		}

		return true
	}

	var doSolve func(curPermutation []int, visitInfo map[int]bool)

	doSolve = func(curPermutation []int, visitInfo map[int]bool) {
		if len(curPermutation) == n {
			result = append(result, generateResultFromPermutation(curPermutation))
			return
		}

		for i := 0; i < n; i++ {
			if !visitInfo[i] {
				nextPermutation := append(curPermutation, i)
				if isValidBoard(nextPermutation) {
					visitInfo[i] = true
					doSolve(nextPermutation, visitInfo)
					visitInfo[i] = false
				}
			}
		}
	}

	doSolve([]int{}, map[int]bool{})

	return result
}

func totalNQueens(n int) int {
	result := 0

	abs := func(v int) int {
		if v < 0 {
			return -v
		}
		return v
	}

	isValidBoard := func(permutation []int) bool {
		for i := 0; i < len(permutation); i++ {
			for j := i + 1; j < len(permutation); j++ {
				if j-i == abs(permutation[i]-permutation[j]) {
					return false
				}
			}
		}

		return true
	}

	var doSolve func(curPermutation []int, visitInfo map[int]bool)

	doSolve = func(curPermutation []int, visitInfo map[int]bool) {
		if len(curPermutation) == n {
			result++
			return
		}

		for i := 0; i < n; i++ {
			if !visitInfo[i] {
				nextPermutation := append(curPermutation, i)
				if isValidBoard(nextPermutation) {
					visitInfo[i] = true
					doSolve(nextPermutation, visitInfo)
					visitInfo[i] = false
				}
			}
		}
	}

	doSolve([]int{}, map[int]bool{})

	return result
}

func maxSubArray(nums []int) int {
	//dp := make([]int, len(nums)) // dp[i] means the max subarray that end with dp[i]
	//max := nums[0]
	//dp[0] = max
	//
	//for i := 1; i < len(nums); i++ {
	//	if nums[i] < 0 {
	//		if dp[i-1] > 0 {
	//			dp[i] = dp[i-1] + nums[i]
	//		} else {
	//			dp[i] = nums[i]
	//		}
	//
	//	} else {
	//		if dp[i-1]+nums[i] <= nums[i] {
	//			dp[i] = nums[i]
	//		} else {
	//			dp[i] = dp[i-1] + nums[i]
	//		}
	//	}
	//	if dp[i] > max {
	//		max = dp[i]
	//	}
	//}
	//
	//return max

	// space save version
	dp := make([]int, 2) // dp[i] means the max subarray that end with dp[i]
	max := nums[0]
	dp[0] = max

	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			if dp[0] > 0 {
				dp[1] = dp[0] + nums[i]
			} else {
				dp[1] = nums[i]
			}

		} else {
			if dp[0]+nums[i] <= nums[i] {
				dp[1] = nums[i]
			} else {
				dp[1] = dp[0] + nums[i]
			}
		}
		if dp[1] > max {
			max = dp[1]
		}
		dp[0] = dp[1]
	}

	return max
}

func spiralOrder(matrix [][]int) []int {
	result := make([]int, len(matrix)*len(matrix[0]))
	index := 0

	var solveSpiralOrder func(matrix [][]int, startX, endX, startY, endY int)

	solveSpiralOrder = func(matrix [][]int, startX, endX, startY, endY int) {
		if startX > endX || startY > endY {
			return
		}
		// travel right
		for j := startY; j <= endY; j++ {
			result[index] = matrix[startX][j]
			index++
		}

		// travel down
		for i := startX + 1; i <= endX; i++ {
			result[index] = matrix[i][endY]
			index++
		}

		if startX != endX {
			// travel left
			for j := endY - 1; j >= startY; j-- {
				result[index] = matrix[endX][j]
				index++
			}
		}

		if startY != endY {
			// travel up
			for i := endX - 1; i >= startX+1; i-- {
				result[index] = matrix[i][startY]
				index++
			}
		}

		// solve next level
		solveSpiralOrder(matrix, startX+1, endX-1, startY+1, endY-1)
	}

	solveSpiralOrder(matrix, 0, len(matrix)-1, 0, len(matrix[0])-1)

	return result
}

func canJump(nums []int) bool {
	furthestJumpCount := 0
	for i, v := range nums {
		furthestJumpCount--
		if v == 0 && furthestJumpCount <= 0 && i != len(nums)-1 {
			return false
		}

		if v > furthestJumpCount {
			furthestJumpCount = v
		}
	}

	return true
}

func merge(intervals [][]int) [][]int {
	//sort.Slice(intervals, func(i, j int) bool {
	//	return intervals[i][0] <= intervals[j][0]
	//})

	writeIndex := 0
	readIndex := 1
	for readIndex < len(intervals) {
		if intervals[readIndex][1] <= intervals[writeIndex][1] {
			readIndex++
		} else if intervals[readIndex][0] <= intervals[writeIndex][1] {
			intervals[writeIndex][1] = intervals[readIndex][1]
			readIndex++
		} else {
			writeIndex++
			intervals[writeIndex] = intervals[readIndex]
			readIndex++
		}
	}

	return intervals[0 : writeIndex+1]
}

func searchInsertPosition(intervals [][]int, v int) int {
	left, right := 0, len(intervals)-1

	for left <= right {
		mid := (left + right) / 2
		if intervals[mid][0] == v {
			return mid
		} else if mid < len(intervals)-1 && v > intervals[mid][0] && v < intervals[mid+1][0] {
			return mid
		} else if mid > 0 && v > intervals[mid-1][0] && v < intervals[mid][0] {
			return mid - 1
		} else if intervals[mid][0] < v {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return right
}

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}
	// binary search insert position

	endInsertPosition := searchInsertPosition(intervals, newInterval[1])
	if endInsertPosition < 0 {
		return append([][]int{newInterval}, intervals...)
	}

	beginInsertPosition := searchInsertPosition(intervals, newInterval[0])

	result := make([][]int, 0)
	for i := 0; i < beginInsertPosition; i++ {
		result = append(result, intervals[i])
	}

	intervalsToMerge := make([][]int, 2, endInsertPosition-beginInsertPosition+2)
	if beginInsertPosition < 0 {
		beginInsertPosition = 0
		intervalsToMerge[0] = newInterval
		intervalsToMerge[1] = intervals[beginInsertPosition]
	} else {
		intervalsToMerge[0] = intervals[beginInsertPosition]
		intervalsToMerge[1] = newInterval
	}

	for i := beginInsertPosition + 1; i <= endInsertPosition; i++ {
		intervalsToMerge = append(intervalsToMerge, intervals[i])
	}
	result = append(result, merge(intervalsToMerge)...)

	for i := endInsertPosition + 1; i < len(intervals); i++ {
		result = append(result, intervals[i])
	}

	return result
}

func lengthOfLastWord(s string) int {
	i := len(s) - 1
	for s[i] == ' ' {
		i--
	}

	j := i
	for j >= 0 && s[j] != ' ' {
		j--
	}
	return i - j
}

func generateMatrix(n int) [][]int {
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
	}

	startX, endX := 0, n-1
	startY, endY := 0, n-1
	v := 1
	for startX <= endX {
		for j := startY; j <= endY; j++ {
			result[startX][j] = v
			v++
		}

		// travel down
		for i := startX + 1; i <= endX; i++ {
			result[i][endY] = v
			v++
		}

		// travel left
		for j := endY - 1; j >= startY; j-- {
			result[endX][j] = v
			v++
		}

		// travel up
		for i := endX - 1; i >= startX+1; i-- {
			result[i][startY] = v
			v++
		}

		startX++
		startY++
		endX--
		endY--
	}

	return result
}

var factorialMap = [9]int{1, 1, 2, 6, 24, 120, 720, 5040, 40320}

func getPermutation(n int, k int) string {
	if n == 1 {
		return "1"
	}
	result := make([]byte, n)
	visitMap := make(map[byte]bool)
	writeIndex := 0
	for writeIndex < n-1 {
		quotient := k / factorialMap[n-1-writeIndex]
		remain := k % factorialMap[n-1-writeIndex]
		if remain != 0 {
			quotient += 1
		} else {
			remain = factorialMap[n-1-writeIndex]
		}
		index := 0
		for i := byte('1'); i < '1'+byte(n); i++ {
			if !visitMap[i] {
				index++
				if index == quotient {
					result[writeIndex] = i
					writeIndex++
					visitMap[i] = true
					break
				}
			}
		}
		k = remain
	}

	for i := byte('1'); i < '1'+byte(n); i++ {
		if !visitMap[i] {
			result[writeIndex] = i
			break
		}
	}

	return string(result)
}
