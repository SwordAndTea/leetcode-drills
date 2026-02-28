package _311_320

import "testing"

func TestShortestDistance(t *testing.T) {
	grid := [][]int{
		{1, 0, 2, 0, 1},
		{0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0},
	}
	t.Log(shortestDistance(grid))
}
