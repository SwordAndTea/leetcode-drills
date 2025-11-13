package _11_220

import "testing"

func Test_WordDictionary(t *testing.T) {
	wordDictionary := &WordDictionary{}
	wordDictionary.AddWord("bad")
	wordDictionary.AddWord("dad")
	wordDictionary.AddWord("mad")
	wordDictionary.Search("pad") // return False
	wordDictionary.Search("bad") // return True
	wordDictionary.Search(".ad") // return True
	wordDictionary.Search("b..") // return True
}

func Test_findKthLargest(t *testing.T) {
	t.Log(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
}

func Test_combinationSum3(t *testing.T) {
	t.Log(combinationSum3(8, 40))
}

func Test_getSkyline(t *testing.T) {
	//t.Log(getSkyline([][]int{
	//	{2, 9, 10},
	//	{3, 7, 15},
	//	{5, 12, 12},
	//	{15, 20, 10},
	//	{19, 24, 8},
	//}))
	t.Log(getSkyline([][]int{
		{0, 2, 3},
		{2, 5, 3},
	}))
	//t.Log(getSkyline([][]int{
	//	{2, 9, 10},
	//	{9, 12, 15},
	//}))
	//t.Log(getSkyline([][]int{
	//	{1, 2, 1},
	//	{1, 2, 2},
	//	{1, 2, 3},
	//}))
}

func Test_containsNearbyDuplicate(t *testing.T) {
	t.Log(containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2))
}

func Test_containsNearbyAlmostDuplicate(t *testing.T) {
	//t.Log(containsNearbyAlmostDuplicate([]int{-3, 2, -6}, 2, 3))
	t.Log(containsNearbyAlmostDuplicate([]int{-1, 1}, 1, 0))
}
