package _2381_2390

import "testing"

func TestLargestPalindromic(t *testing.T) {
	if largestPalindromic("444947137") != "7449447" {
		t.FailNow()
	}
}
