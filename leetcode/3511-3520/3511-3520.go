package _3511_3520

import (
	"slices"
)

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

// leetcode problem No. 3518
func smallestPalindrome2(s string, k int) string {
	strCount := map[rune]int{}
	for _, ch := range s {
		strCount[ch]++
	}

	var backtracking func(curPermutation []rune)
	visit := make(map[rune]bool)
	curK := 0
	var kthPermuation string
	targetLen := len(strCount)
	var mid rune
	for ch, count := range strCount {
		if count == 1 {
			targetLen -= 1
		}
		if count%2 == 1 {
			mid = ch
		}
	}
	backtracking = func(curPermutation []rune) {
		if curK > k {
			return
		}
		if len(curPermutation) == targetLen {
			curK++
			if curK == k {
				kthPermuation = string(curPermutation)
			}
			return
		}

		for i := 'a'; i <= 'z'; i++ {
			if strCount[i] > 1 && !visit[i] { // note: we use > 1 not > 0, as if the count is 1, it can only be put at the middle
				visit[i] = true
				backtracking(append(curPermutation, i))
				visit[i] = false
			}
		}
	}

	backtracking(nil)
	if curK < k {
		return ""
	}

	frontHalf := make([]rune, 0, len(s)/2)

	for _, ch := range kthPermuation {
		for i := 0; i < strCount[ch]/2; i++ {
			frontHalf = append(frontHalf, ch)
		}
	}
	backHalf := make([]rune, len(frontHalf))
	copy(backHalf, frontHalf)
	slices.Reverse(backHalf)
	if mid != 0 {
		return string(frontHalf) + string(mid) + string(backHalf)
	}
	return string(frontHalf) + string(backHalf)
}
