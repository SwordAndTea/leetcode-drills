package _31_40

import (
	"sort"
	"strconv"
)

func nextPermutation(nums []int) {
	if len(nums) == 1 {
		return
	}
	if len(nums) == 2 {
		nums[0], nums[1] = nums[1], nums[0]
	}

	// find the beginning of last ascending sequence
	i := len(nums) - 1
	for ; i > 0 && nums[i-1] >= nums[i]; i-- {
	}

	if i == 0 {
		// reverse
		for m, n := 0, len(nums)-1; m < n; {
			nums[m], nums[n] = nums[n], nums[m]
			m++
			n--
		}
		return
	}

	j := i - 1

	for i = len(nums) - 1; i > j; i-- {
		if nums[i] > nums[j] {
			break
		}
	}

	nums[i], nums[j] = nums[j], nums[i]
	// reverse nums[j + 1: len(nums)]
	for m, n := j+1, len(nums)-1; m < n; {
		nums[m], nums[n] = nums[n], nums[m]
		m++
		n--
	}
}

// leetcode problem No. 32

func longestValidParentheses(s string) int {
	n := len(s)
	// dp[i] stands for longest ValidParentheses that ends the s[i]
	dp := make([]int, n)

	ans := 0
	for i := 1; i < n; i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				if i-2 >= 0 {
					dp[i] = dp[i-2] + 2
				}
			} else {
				j := i - 1 - dp[i-1]
				if j >= 0 && s[j] == '(' {
					if j-1 >= 0 {
						dp[i] = dp[j-1] + dp[i-1] + 2
					} else {
						dp[i] = dp[i-1] + 2
					}
				}
			}
			ans = max(ans, dp[i])
		}
	}

	return ans
}

// Search in Rotated Sorted Array

func findRotatePivot(nums []int) int {
	left, right := 0, len(nums)-1
	mid := 0
	for left < right {
		mid = (left + right) / 2
		if mid <= len(nums)-2 && nums[mid] > nums[mid+1] {
			return mid
		}
		if mid >= 1 && nums[mid-1] > nums[mid] {
			return mid - 1
		}
		if nums[mid] >= nums[right] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func binarySearch(nums []int, target int, left, right int) int {
	mid := 0
	for left <= right {
		mid = (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func search(nums []int, target int) int {
	pivot := findRotatePivot(nums)

	if pivot == -1 {
		return binarySearch(nums, target, 0, len(nums)-1)
	}

	if target <= nums[len(nums)-1] {
		return binarySearch(nums, target, pivot+1, len(nums)-1)
	}

	return binarySearch(nums, target, 0, pivot)
}

// Find First and Last Position of Element in Sorted Array
func searchRange(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	mid := 0
	for left <= right {
		mid = (left + right) / 2
		if nums[mid] == target {
			// find a target, starting find left and right
			targetLeft, targetRight := 0, len(nums)-1
			left, right = 0, mid
			newMid := 0
			for left <= right {
				newMid = (left + right) / 2
				if newMid > 0 && nums[newMid] == target && nums[newMid-1] != target {
					targetLeft = newMid
					break
				}
				if nums[newMid] < target {
					left = newMid + 1
				} else {
					right = newMid - 1
				}
			}

			left, right = mid, len(nums)-1

			for left <= right {
				newMid = (left + right) / 2
				if newMid < len(nums)-1 && nums[newMid] == target && nums[newMid+1] != target {
					targetRight = newMid
					break
				}

				if nums[newMid] > target {
					right = newMid - 1
				} else {
					left = newMid + 1
				}
			}

			return []int{targetLeft, targetRight}
		}
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return []int{-1, -1}
}

// search insert position
func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	mid := 0
	for left <= right {
		mid = (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

func isValidSudoku(board [][]byte) bool {
	var rowNumInfo [9]int
	var colNumInfo [9]int
	var boxNumInfo [9]int

	boxIndex := func(i, j int) int {
		return (i/3)*3 + j/3
	}

	canMaskBeSet := func(i, j, mask int) bool {
		if rowNumInfo[i]&mask > 0 {
			return false
		}
		if colNumInfo[j]&mask > 0 {
			return false
		}
		if boxNumInfo[boxIndex(i, j)]&mask > 0 {
			return false
		}
		return true
	}

	setMask := func(i, j, mask int) {
		rowNumInfo[i] |= mask
		colNumInfo[j] |= mask
		boxNumInfo[boxIndex(i, j)] |= mask
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				mask := 1 << (board[i][j] - '1')
				if !canMaskBeSet(i, j, mask) {
					return false
				}

				setMask(i, j, mask)
			}
		}
	}

	return true
}

func solveSudoku(board [][]byte) {
	var rowNumInfo [9]int
	var colNumInfo [9]int
	var boxNumInfo [9]int

	boxIndex := func(i, j int) int {
		return (i/3)*3 + j/3
	}

	canMaskBeSet := func(i, j, mask int) bool {
		if rowNumInfo[i]&mask > 0 {
			return false
		}
		if colNumInfo[j]&mask > 0 {
			return false
		}
		if boxNumInfo[boxIndex(i, j)]&mask > 0 {
			return false
		}
		return true
	}

	setMask := func(i, j, mask int) {
		rowNumInfo[i] |= mask
		colNumInfo[j] |= mask
		boxNumInfo[boxIndex(i, j)] |= mask
	}

	unsetMask := func(i, j, mask int) {
		rowNumInfo[i] ^= mask
		colNumInfo[j] ^= mask
		boxNumInfo[boxIndex(i, j)] |= mask
	}

	emptyPositions := make([][2]int, 0, 32)

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				emptyPositions = append(emptyPositions, [2]int{i, j})
			} else {
				mask := 1 << (board[i][j] - '1')
				setMask(i, j, mask)
			}
		}
	}

	var doSolvedSudoku func(i int) bool

	doSolvedSudoku = func(i int) bool {
		if i == len(emptyPositions) {
			return true
		}

		pos := emptyPositions[i]
		for j := 0; j < 9; j++ {
			newMask := 1 << j
			if canMaskBeSet(pos[0], pos[1], newMask) {
				setMask(pos[0], pos[1], newMask)
				if doSolvedSudoku(i + 1) {
					board[pos[0]][pos[1]] = byte('1' + j)
					return true
				}
				unsetMask(pos[0], pos[1], newMask)
			}
		}
		return false
	}

	doSolvedSudoku(0)
}

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	strToSay := countAndSay(n - 1)
	i := 0
	count := 0
	newStr := make([]byte, 0)
	for i < len(strToSay) {
		count = 1
		c := strToSay[i]
		for i+1 < len(strToSay) && strToSay[i+1] == strToSay[i] {
			count++
			i++
		}
		newStr = append(append(newStr, []byte(strconv.Itoa(count))...), c)
		i++
	}
	return string(newStr)
}

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	result := make([][]int, 0)

	var solveCombinationSum func(currentCandidates []int, sum int, startIndex int)
	solveCombinationSum = func(currentCandidates []int, sum int, startIndex int) {
		for i := startIndex; i < len(candidates); i++ {
			v := candidates[i]
			if sum+v == target {
				newRes := make([]int, len(currentCandidates)+1)
				copy(newRes, append(currentCandidates, v))
				result = append(result, newRes)
				return
			}

			if sum+v > target {
				return
			}

			solveCombinationSum(append(currentCandidates, v), sum+v, i)
		}
	}

	solveCombinationSum([]int{}, 0, 0)

	return result
}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	result := make([][]int, 0)

	var solveCombinationSum func(currentCandidates []int, sum int, startIndex int)
	solveCombinationSum = func(currentCandidates []int, sum int, startIndex int) {
		for i := startIndex; i < len(candidates); i++ {
			v := candidates[i]
			if sum+v == target {
				newRes := make([]int, len(currentCandidates)+1)
				copy(newRes, append(currentCandidates, v))
				result = append(result, newRes)
				return
			}

			if sum+v > target {
				return
			}

			solveCombinationSum(append(currentCandidates, v), sum+v, i+1)
			for i+1 < len(candidates) && candidates[i+1] == candidates[i] {
				i++
			}
		}
	}

	solveCombinationSum([]int{}, 0, 0)

	return result
}
