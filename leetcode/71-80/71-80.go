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
	len1, len2 := len(word1), len(word2)
	dp := make([][]int, 2)
	dp[0] = make([]int, len2+1)
	dp[1] = make([]int, len2+1)

	for j := 0; j <= len2; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= len1; i++ {
		dp[1][0] = i
		for j := 1; j <= len2; j++ {
			if word1[i-1] == word2[j-1] {
				dp[1][j] = dp[0][j-1]
			} else {
				dp[1][j] = dp[0][j-1] + 1 // replace
				if dp[0][j]+1 < dp[1][j] {
					dp[1][j] = dp[0][j] + 1 // delete
				}

				if dp[1][j-1]+1 < dp[1][j] { // insert
					dp[1][j] = dp[1][j-1] + 1
				}
			}
		}
		dp[0], dp[1] = dp[1], dp[0]
	}

	return dp[0][len2]
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

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	left, right := 0, m-1

	row := -2
	for left <= right {
		mid := (left + right) / 2
		if matrix[mid][0] == target {
			return true
		}
		if mid < m-1 && matrix[mid][0] < target && matrix[mid+1][0] > target {
			row = mid
			break
		}

		if mid > 1 && matrix[mid-1][0] < target && matrix[mid][0] > target {
			row = mid - 1
			break
		}

		if matrix[mid][0] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	if row == -2 {
		row = right
	}

	if row == -1 {
		return false
	}

	if row == m-1 && matrix[row][n-1] < target {
		return false
	}

	left, right = 0, n-1
	for left <= right {
		mid := (left + right) / 2
		if matrix[row][mid] == target {
			return true
		}

		if matrix[row][mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
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

func combine(n int, k int) [][]int {
	combinationValue := func(n, m int) int {
		res := 1
		for i := 1; i <= m; i++ {
			res = res * (n - m + i) / i
		}
		return res
	}

	result := make([][]int, 0, combinationValue(n, k))

	var doCombine func(start int, current []int)

	doCombine = func(start int, current []int) {
		if len(current) == k {
			newRes := make([]int, k)
			copy(newRes, current)
			result = append(result, newRes)
			return
		}
		for i := start; i <= n; i++ {
			doCombine(i+1, append(current, i))
		}
	}

	doCombine(1, []int{})

	return result
}

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
