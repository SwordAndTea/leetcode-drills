package _1_70

import (
	"errors"
	"math"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	if k == 0 {
		return head
	}
	p := head
	length := 0
	var end *ListNode
	for p != nil {
		length++
		if p.Next == nil {
			end = p
		}
		p = p.Next
	}
	k = k % length
	if k == 0 {
		return head
	}
	p = head
	i := 1
	for i < length-k {
		p = p.Next
		i++
	}
	end.Next = head
	head = p.Next
	p.Next = nil
	return head
}

func uniquePaths(m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	}
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			result[j] += result[j-1]
		}
	}
	return result[n-1]
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	result := make([][]int, m)
	result[0] = make([]int, n)
	result[0][0] = 1
	for i := 1; i < m; i++ {
		result[i] = make([]int, n)
		if obstacleGrid[i][0] == 1 {
			result[i][0] = 0
		} else {
			result[i][0] = result[i-1][0]
		}
	}
	for j := 1; j < n; j++ {
		if obstacleGrid[0][j] == 1 {
			result[0][j] = 0
		} else {
			result[0][j] = result[0][j-1]
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				result[i][j] = 0
			} else {
				result[i][j] = result[i-1][j] + result[i][j-1]
			}

		}
	}
	return result[m-1][n-1]
}

func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	result := make([][]int, m)
	result[0] = make([]int, n)
	result[0][0] = grid[0][0]
	for i := 1; i < m; i++ {
		result[i] = make([]int, n)
		result[i][0] = grid[i][0] + result[i-1][0]
	}
	for j := 1; j < n; j++ {
		result[0][j] = grid[0][j] + result[0][j-1]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if result[i-1][j] <= result[i][j-1] {
				result[i][j] = grid[i][j] + result[i-1][j]
			} else {
				result[i][j] = grid[i][j] + result[i][j-1]
			}
		}
	}
	return result[m-1][n-1]
}

func isNumber(s string) bool {
	v, err := strconv.ParseFloat(s, 64)
	if errors.Is(err, strconv.ErrSyntax) {
		return false
	}
	if err == nil && (math.IsInf(v, 1) || math.IsInf(v, -1) || math.IsNaN(v)) {
		return false
	}
	return true
}

func plusOne(digits []int) []int {
	carry := 0
	digits[len(digits)-1] += 1
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] += carry
		carry = digits[i] / 10
		if carry == 0 {
			return digits
		}
		digits[i] = digits[i] % 10
	}
	return append([]int{1}, digits...)
}

func addBinary(a string, b string) string {
	lenA, lenB := len(a), len(b)
	max := lenA
	if lenB > lenA {
		max = lenB
	}
	result := make([]byte, max+1)
	carry := uint8(0)
	i, j := lenA-1, lenB-1
	index := max
	for i >= 0 && j >= 0 {
		result[index] = a[i] + b[j] + carry - '0'*2
		carry = result[index] / 2
		result[index] = '0' + (result[index] % 2)
		i--
		j--
		index--
	}
	for i >= 0 {
		result[index] = a[i] + carry - '0'
		carry = result[index] / 2
		result[index] = '0' + (result[index] % 2)
		i--
		index--
	}
	for j >= 0 {
		result[index] = b[j] + carry - '0'
		carry = result[index] / 2
		result[index] = '0' + (result[index] % 2)
		j--
		index--
	}
	if carry == 0 {
		return string(result[1:])
	}
	result[0] = '1'
	return string(result)
}

func fullJustify(words []string, maxWidth int) []string {
	i := 0
	result := make([]string, 0, 2)
	for i < len(words) {
		curWordLen := len(words[i])
		j := i + 1
		spaceCount := 0
		for j < len(words) {
			curWordLen += len(words[j]) + 1
			if curWordLen > maxWidth {
				curWordLen -= len(words[j]) + 1
				break
			}
			spaceCount++
			j++
		}

		line := make([]byte, maxWidth)
		if j == len(words) { // last line
			copy(line[0:len(words[i])], words[i])
			index := len(words[i])
			for k := i + 1; k < j; k++ {
				line[index] = ' '
				index++
				copy(line[index:index+len(words[k])], words[k])
				index += len(words[k])
			}
			for index < maxWidth {
				line[index] = ' '
				index++
			}
			result = append(result, string(line))
		} else {
			actualWorldLen := curWordLen - spaceCount
			totalSpace := maxWidth - actualWorldLen
			if spaceCount == 0 {
				copy(line[0:len(words[i])], words[i])
				index := len(words[i])
				for index < maxWidth {
					line[index] = ' '
					index++
				}
			} else {
				remain := totalSpace % spaceCount
				quotient := totalSpace / spaceCount
				copy(line[0:len(words[i])], words[i])
				index := len(words[i])
				for k := i + 1; k < j; k++ {
					if remain != 0 {
						for m := 0; m < quotient+1; m++ {
							line[index] = ' '
							index++
						}
						remain--
					} else {
						for m := 0; m < quotient; m++ {
							line[index] = ' '
							index++
						}
					}

					copy(line[index:index+len(words[k])], words[k])
					index += len(words[k])
				}
			}

			result = append(result, string(line))
		}
		i = j
	}
	return result
}

func mySqrt(x int) int {
	base := 1
	v := 1
	for v < x {
		v *= 4
		base *= 2
	}

	if v == x {
		return base
	}

	left := base / 2
	right := base
	for left <= right {
		mid := (left + right) / 2
		if mid*mid == x || (mid*mid > x && (mid+1)*(mid+1) < x) {
			return mid
		}
		if mid*mid > x {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return right
}

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	result := make([]int, n)
	result[0] = 1
	result[1] = 2
	for i := 2; i < n; i++ {
		result[i] = result[i-1] + result[i-2]
	}
	return result[n-1]
}
