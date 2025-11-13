package dp

import (
	"testing"
)

func TestLargestSequenceSum(t *testing.T) {
	sum := LargestSequenceSum([]int{-2, 11, -4, 13, -5, -2})
	t.Log(sum)
}

func TestLongestIncreasingSequence(t *testing.T) {
	length := LongestIncreasingSequence([]int{1, 2, 3, -9, 3, 9, 0, 11})
	t.Log(length)
}
