package weekly_contest

import "testing"

func Test_numberOfSpecialChars(t *testing.T) {
	t.Log(numberOfSpecialChars("aaAAbcBC"))
	t.Log(numberOfSpecialChars("abc"))
	t.Log(numberOfSpecialChars("abBCab"))
}

func Test_numberOfSpecialChars2(t *testing.T) {
	t.Log(numberOfSpecialChars2("aaAAbcBC"))
	t.Log(numberOfSpecialChars2("abc"))
	t.Log(numberOfSpecialChars2("abBCab"))
}

func Test_minimumOperations(t *testing.T) {
	t.Log(minimumOperations([][]int{
		{1, 6, 7, 3, 0, 4, 1, 3, 7, 5},
	}))
}
