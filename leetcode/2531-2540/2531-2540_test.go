package _2531_2540

import "testing"

func TestTimeTaken(t *testing.T) {
	t.Logf("%+v", timeTaken([]int{0, 1, 1, 2, 4}, []int{0, 1, 0, 0, 1}))
	t.Logf("%+v", timeTaken([]int{0, 0, 0}, []int{1, 0, 1}))
}

func TestRangeAddQueries(t *testing.T) {
	t.Logf("%#v", rangeAddQueries(3, [][]int{
		{1, 1, 2, 2},
		{0, 0, 1, 1},
	}))
}
