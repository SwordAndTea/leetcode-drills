package _2501_2510

import "testing"

func TestAllocator(t *testing.T) {
	loc := Constructor(7)
	t.Log(loc.Allocate(7, 8))
	t.Log(loc.FreeMemory(8))
	t.Log(loc.Allocate(7, 6))
}
