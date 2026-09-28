package _1181_1190

import "slices"

// leetcode problem No. 1190
func reverseParentheses(s string) string {
	strByte := []byte(s)

	var leftIndexStack []int
	for i := 0; i < len(strByte); i++ {
		if strByte[i] == '(' {
			leftIndexStack = append(leftIndexStack, i)
		} else if strByte[i] == ')' {
			leftIndex := leftIndexStack[len(leftIndexStack)-1]
			leftIndexStack = leftIndexStack[:len(leftIndexStack)-1]
			slices.Reverse(strByte[leftIndex+1 : i])
		}
	}
	ans := make([]byte, 0, len(strByte))
	for _, c := range strByte {
		if c == '(' || c == ')' {
			continue
		}
		ans = append(ans, c)
	}
	return string(ans)
}
