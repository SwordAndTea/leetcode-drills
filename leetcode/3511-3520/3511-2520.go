package _3511_3520

import "slices"

// leetcode problem No. 3517
func smallestPalindrome(s string) string {
	strCount := map[rune]int{}
	for _, ch := range s {
		strCount[ch]++
	}
	front := make([]rune, 0, len(s))
	var midCh rune
	for i := 'a'; i <= 'z'; i++ {
		for j := 0; j < strCount[i]/2; j++ {
			front = append(front, i)
		}
		if strCount[i]%2 == 1 {
			midCh = i
		}
	}
	back := make([]rune, len(front))
	copy(back, front)
	slices.Reverse(back)

	if midCh != 0 {
		return string(front) + string(midCh) + string(back)
	}
	return string(front) + string(back)
}
