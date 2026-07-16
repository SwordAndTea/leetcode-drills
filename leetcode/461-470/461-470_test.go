package _461_470

import "testing"

func TestIslandPerimeter(t *testing.T) {
	grid := [][]int{
		{0, 1, 0, 0},
		{1, 1, 1, 0},
		{0, 1, 0, 0},
		{1, 1, 0, 0},
	}
	t.Log(islandPerimeter(grid))
}
