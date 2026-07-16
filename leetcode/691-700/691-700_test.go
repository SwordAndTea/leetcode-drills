package _691_700

import "testing"

func TestNumDistinctIslands(t *testing.T) {
	t.Log(numDistinctIslands([][]int{
		{1, 1, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 0, 1, 1},
		{0, 0, 0, 1, 1},
	}))
	t.Log(numDistinctIslands([][]int{
		{1, 1, 0, 1, 1},
		{1, 0, 0, 0, 0},
		{0, 0, 0, 0, 1},
		{1, 1, 0, 1, 1},
	}))
}

func TestTopKFrequent(t *testing.T) {
	t.Log(topKFrequent([]string{"i", "love", "leetcode", "i", "love", "coding"}, 2))
}

func TestFindShortestSubArray(t *testing.T) {
	t.Log(findShortestSubArray([]int{1, 2, 2, 1, 2, 1, 1, 1, 1, 2, 2, 2}))
}

func TestCanPartitionKSubsets(t *testing.T) {
	t.Log(canPartitionKSubsets([]int{4, 3, 2, 3, 5, 2, 1}, 4))
}
