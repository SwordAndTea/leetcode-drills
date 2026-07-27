package _2531_2540

import "testing"

func TestRangeAddQueries(t *testing.T) {
	t.Logf("%#v", rangeAddQueries(3, [][]int{
		{1, 1, 2, 2},
		{0, 0, 1, 1},
	}))
}
