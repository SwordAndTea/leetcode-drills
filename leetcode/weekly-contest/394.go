package weekly_contest

import "math"

func numberOfSpecialChars(word string) int {
	result := 0
	letterMap := make(map[uint8]bool)
	for i := 0; i < len(word); i++ {
		c := word[i]
		if c >= 'a' && c <= 'z' {
			if !letterMap[c] && letterMap[c-32] {
				result++
			}
		} else {
			if !letterMap[c] && letterMap[c+32] {
				result++
			}
		}
		letterMap[c] = true
	}
	return result
}

func numberOfSpecialChars2(word string) int {
	letterMap := make(map[uint8]bool)
	for i := 0; i < len(word); i++ {
		c := word[i]
		if c >= 'a' && c <= 'z' {
			if letterMap[c-32] { // if upper has appeared before
				letterMap[c] = false
			} else {
				letterMap[c] = true
			}
		} else {
			letterMap[c] = true
		}
	}
	result := 0
	for i := uint8('a'); i <= 'z'; i++ {
		if letterMap[i] && letterMap[i-32] {
			result++
		}
	}
	return result
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func minimumOperations(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	numberMap := make([][10]int, n)
	// dp[i][j] means the min op count that i count need when column i-1's value is j
	dp := make([][10]int, n)

	for j := 0; j < n; j++ {
		for i := 0; i < m; i++ {
			numberMap[j][grid[i][j]]++
		}
		for i := 0; i < 10; i++ {
			dp[j][i] = math.MaxInt
		}
	}

	var dfsDP func(curJ, preReplaceValue int) int

	dfsDP = func(curJ, preReplaceValue int) int {
		if curJ == n {
			return 0
		}

		if dp[curJ][preReplaceValue] == math.MaxInt {
			for v := 0; v < 10; v++ {
				if curJ == 0 || v != preReplaceValue {
					count := m - numberMap[curJ][v] + dfsDP(curJ+1, v)
					dp[curJ][preReplaceValue] = min(dp[curJ][preReplaceValue], count)
				}
			}
		}

		return dp[curJ][preReplaceValue]
	}

	return dfsDP(0, 0)
}
