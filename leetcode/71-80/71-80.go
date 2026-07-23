package _71_70

import (
	"math"
	"strings"
)

func simplifyPath(path string) string {
	if len(path) == 1 {
		return path
	}
	index := 1
	var onePath string
	pathList := make([]string, 0, 4)
	for index < len(path) {
		i := index
		for ; i < len(path); i++ {
			if path[i] == '/' {
				break
			}
		}
		onePath = path[index:i]
		if onePath != "" {
			if onePath == ".." {
				if len(pathList) > 0 {
					pathList = pathList[0 : len(pathList)-1]
				}
			} else if onePath != "." {
				pathList = append(pathList, onePath)
			}
		}

		for ; i < len(path) && path[i] == '/'; i++ {

		}
		index = i
	}
	return strings.Join([]string{"/", strings.Join(pathList, "/")}, "")
}

// leetcode problem No. 72
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}
	for j := 1; j <= n; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i][j-1], dp[i-1][j], dp[i-1][j-1]) + 1
			}
		}
	}
	return dp[m][n]
}

func setZeroes(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])
	zeroX := make(map[int]bool)
	zeroY := make(map[int]bool)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				zeroX[i] = true
				zeroY[j] = true
			}
		}
	}

	for k, _ := range zeroX {
		for j := 0; j < n; j++ {
			matrix[k][j] = 0
		}
	}

	for k, _ := range zeroY {
		for i := 0; i < m; i++ {
			matrix[i][k] = 0
		}
	}
}

// leetcode problem No. 74
func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	rowLeft, rowRight := 0, m-1
	colLeft, colRight := 0, n-1
	for rowLeft <= rowRight && colLeft <= colRight {
		midRow := (rowLeft + rowRight) / 2
		midCol := (colLeft + colRight) / 2
		if matrix[midRow][midCol] == target {
			return true
		}
		if target > matrix[midRow][n-1] {
			rowLeft = midRow + 1
		} else if target < matrix[midRow][0] {
			rowRight = midRow - 1
		} else { // matrix[midRow][0] <= target <= matrix[midRow][n-1], we located the row
			if target > matrix[midRow][midCol] {
				colLeft = midCol + 1
			} else {
				colRight = midCol - 1
			}
		}
	}

	return false
}

func sortColors(nums []int) {
	//partition := func(start, end int) int {
	//	guard := nums[start]
	//	left, right := start, end
	//	for left < right {
	//		for left < right && nums[right] >= guard {
	//			right--
	//		}
	//
	//		for left < right && nums[left] <= guard {
	//			left++
	//		}
	//		nums[left], nums[right] = nums[right], nums[left]
	//	}
	//	nums[left], nums[start] = nums[start], nums[left]
	//	return left
	//}
	//
	//var quickSort func(start, end int)
	//
	//quickSort = func(start, end int) {
	//	if start < end {
	//		partitionIndex := partition(start, end)
	//		quickSort(start, partitionIndex-1)
	//		quickSort(partitionIndex+1, end)
	//	}
	//}
	//
	//quickSort(0, len(nums)-1)
	numCount := make(map[int]int)
	for _, v := range nums {
		numCount[v] += 1
	}
	for i := 0; i < len(nums); i++ {
		if i < numCount[0] {
			nums[i] = 0
		} else if i < numCount[0]+numCount[1] {
			nums[i] = 1
		} else {
			nums[i] = 2
		}
	}
}

// leetcode problem No. 76
func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	charCounts := make(map[uint8]int)
	for _, c := range t {
		charCounts[uint8(c)] += 1
	}

	matchCount := 0
	matchInfo := make(map[uint8]int)
	left, right := 0, 0
	minLen := math.MaxInt
	minString := ""
	for right < len(s) {
		c := s[right]
		if charCounts[c] != 0 {
			if matchInfo[c] < charCounts[c] {
				matchCount++
			}
			matchInfo[c]++
		}
		right++
		if matchCount == len(t) {
			// move left
			for left < right {
				c = s[left]
				if charCounts[c] != 0 {
					matchInfo[c]--
					if matchInfo[c] < charCounts[c] {
						if right-left < minLen {
							minLen = right - left
							minString = s[left:right]
						}
						left++
						matchCount--
						break
					}
				}
				left++
			}
		}
	}
	if minLen == math.MaxInt {
		return ""
	}
	return minString
}

// leetcode problem No. 77
func combine(n int, k int) [][]int {
	result := make([][]int, 0)
	var backtracking func(int, []int)
	backtracking = func(start int, curNums []int) {
		if len(curNums) == k {
			tmp := make([]int, k)
			copy(tmp, curNums)
			result = append(result, tmp)
			return
		}
		for i := start + 1; i <= n; i++ {
			backtracking(i, append(curNums, i))
		}
	}

	backtracking(0, []int{})
	return result
}

// leetcode problem No. 78
func subsets(nums []int) [][]int {
	result := make([][]int, 2)
	result[0] = []int{}
	result[1] = nums

	var getSubset func(startIndex int, current []int)

	getSubset = func(startIndex int, current []int) {
		for i := startIndex; i < len(nums); i++ {
			if len(current)+1 < len(nums) {
				newResult := make([]int, len(current)+1)
				copy(newResult, current)
				newResult[len(current)] = nums[i]
				result = append(result, newResult)
				getSubset(i+1, append(current, nums[i]))
			}
		}
	}

	getSubset(0, []int{})

	return result
}

func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	if m*n < len(word) {
		return false
	}

	var depthSearch func(curI, curJ, curWordIndex int, visitInfo [][]bool) bool

	depthSearch = func(curI, curJ, curWordIndex int, visitInfo [][]bool) bool {
		if curWordIndex == len(word) {
			return true
		}
		if curJ-1 >= 0 && !visitInfo[curI][curJ-1] && board[curI][curJ-1] == word[curWordIndex] { // go left
			visitInfo[curI][curJ-1] = true
			if depthSearch(curI, curJ-1, curWordIndex+1, visitInfo) {
				return true
			}
			visitInfo[curI][curJ-1] = false
		}

		if curJ+1 < n && !visitInfo[curI][curJ+1] && board[curI][curJ+1] == word[curWordIndex] { // go right
			visitInfo[curI][curJ+1] = true
			if depthSearch(curI, curJ+1, curWordIndex+1, visitInfo) {
				return true
			}
			visitInfo[curI][curJ+1] = false
		}

		if curI-1 >= 0 && !visitInfo[curI-1][curJ] && board[curI-1][curJ] == word[curWordIndex] { // go up
			visitInfo[curI-1][curJ] = true
			if depthSearch(curI-1, curJ, curWordIndex+1, visitInfo) {
				return true
			}
			visitInfo[curI-1][curJ] = false
		}

		if curI+1 < m && !visitInfo[curI+1][curJ] && board[curI+1][curJ] == word[curWordIndex] { // go up
			visitInfo[curI+1][curJ] = true
			if depthSearch(curI+1, curJ, curWordIndex+1, visitInfo) {
				return true
			}
			visitInfo[curI+1][curJ] = false
		}

		return false
	}

	visitInfo := make([][]bool, m)

	for i := 0; i < m; i++ {
		visitInfo[i] = make([]bool, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[0] {
				visitInfo[i][j] = true
				if depthSearch(i, j, 1, visitInfo) {
					return true
				}
				visitInfo[i][j] = false
			}
		}
	}

	return false
}

func removeDuplicates(nums []int) int {
	readIndex, writeIndex := 1, 1
	pre := nums[0]
	count := 1
	for readIndex < len(nums) {
		if nums[readIndex] == pre {
			count += 1
		} else {
			count = 1
		}

		if count <= 2 {
			nums[writeIndex] = nums[readIndex]
			writeIndex++
		}

		pre = nums[readIndex]
		readIndex++
	}
	return writeIndex
}
