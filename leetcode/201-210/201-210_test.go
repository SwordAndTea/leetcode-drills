package _201_210

import "testing"

func Test_isIsomorphic(t *testing.T) {
	t.Log(isIsomorphic("paper", "title"))
}

func Test_canFinish(t *testing.T) {
	t.Log(canFinish(2, [][]int{{1, 0}}))
}

func Test_Trie(t *testing.T) {
	trie := Constructor()
	trie.Insert("apple")
	t.Log(trie.Search("apple"))
	t.Log(trie.Search("app"))
	t.Log(trie.StartsWith("app"))
}
