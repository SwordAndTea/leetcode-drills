package _1961_1970

import "strings"

// leetcode problem No. 1967
func numOfStrings(patterns []string, word string) int {
	result := 0
	for _, pattern := range patterns {
		if strings.Contains(word, pattern) {
			result++
		}
	}
	return result
}
