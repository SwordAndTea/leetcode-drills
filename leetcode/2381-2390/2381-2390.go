package _2381_2390

import "strings"

func reverseString(value string) string {
	result := make([]rune, len(value))
	for i, v := range value {
		result[len(value)-1-i] = v
	}
	return string(result)
}

// leetcode problem No. 2384
func largestPalindromic(num string) string {
	if len(num) == 1 {
		return num
	}
	numCount := make(map[rune]int)
	for _, v := range num {
		numCount[v]++
	}

	if len(numCount) == 1 && numCount['0'] != 0 {
		return "0"
	}

	largestNumWithOddNum := rune(0)
	for i := '9'; i >= '0'; i-- {
		v := numCount[i]
		if v != 0 && v%2 == 1 {
			largestNumWithOddNum = i
			break
		}
	}

	result := ""
	// construct first half string
	for i := '9'; i >= '0'; i-- {
		v := numCount[i]
		if i == '0' && result == "" {
			continue
		}
		if v != 0 {
			result += strings.Repeat(string(i), v/2)
		}
	}
	if largestNumWithOddNum != rune(0) {
		result = result + string(largestNumWithOddNum) + reverseString(result)
	} else {
		result += reverseString(result)
	}
	return result
}
