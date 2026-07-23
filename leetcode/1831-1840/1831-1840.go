package _1831_1840

// leetcode problem No. 1839
var vowelOrder = map[rune]int{
	'a': 1,
	'e': 2,
	'i': 3,
	'o': 4,
	'u': 5,
}

func longestBeautifulSubstring(word string) int {
	left := -1
	ans := 0
	lastOrder := 0
	for i, c := range word {
		if order, ok := vowelOrder[c]; ok && (order == lastOrder || order-lastOrder == 1) {
			if order == 5 {
				ans = max(ans, i-left)
			}
			lastOrder = order
		} else {
			if c == 'a' {
				left = i - 1
				lastOrder = 1
			} else {
				left = i
				lastOrder = 0
			}
		}
	}
	return ans
}
