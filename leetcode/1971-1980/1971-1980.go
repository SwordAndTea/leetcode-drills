package _1971_1980

import "math"

// leetcode problem No. 1975
func maxMatrixSum(matrix [][]int) int {
	sum := 0
	minAbs := math.MaxInt
	n := len(matrix)
	numNegative := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			absVal := matrix[i][j]
			if matrix[i][j] < 0 {
				numNegative++
				absVal = -absVal
			}
			if absVal < minAbs {
				minAbs = absVal
			}
			sum += absVal
		}
	}
	if numNegative%2 == 0 {
		return sum
	}
	return sum - minAbs*2
}
