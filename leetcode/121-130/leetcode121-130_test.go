package _21_130

import "testing"

func Test_maxProfit2(t *testing.T) {
	t.Log(maxProfit3([]int{1, 2, 3, 4, 5}))
}

func Test_isPalindrome(t *testing.T) {
	t.Log(isPalindrome("A man, a plan, a canal: Panama"))
}

func Test_findLadders(t *testing.T) {
	t.Log(findLadders("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}))
}

func Test_ladderLength(t *testing.T) {
	t.Log(ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}))
}

func Test_solve(t *testing.T) {
	board := [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', '0', '0', 'X'},
		{'X', 'X', '0', 'X'},
		{'X', '0', 'X', 'X'},
	}
	t.Log(board)
	solve(board)
	t.Log(board)
}
