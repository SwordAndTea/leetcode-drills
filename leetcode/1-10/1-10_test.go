package _1_10

import "testing"

func TestModulo(t *testing.T) {
	t.Log(-23 % 10)
	t.Log(-23 / 10)
	t.Log(-2 / 10)
}

func TestIsMatch(t *testing.T) {
	t.Log(isMatch("aa", "a*"))
}

func TestLengthOfLongestSubstring(t *testing.T) {
	t.Log(lengthOfLongestSubstring("abcdd"))
}
