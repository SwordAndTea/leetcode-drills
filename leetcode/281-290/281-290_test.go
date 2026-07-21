package _281_290

import "testing"

func TestGameOfLife(t *testing.T) {
	board := [][]int{
		{0, 1, 0},
		{0, 0, 1},
		{1, 1, 1},
		{0, 0, 0},
	}
	gameOfLife(board)
	t.Logf("%#v", board)
}
