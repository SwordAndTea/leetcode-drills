package _391_400

import (
	"strings"
)

// leetcode problem No. 394
func decodeString(s string) string {
	numStack := []int{}
	strStack := []string{}
	ans := ""
	curNum := 0

	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			curNum = curNum*10 + int(s[i]-'0')
		} else if s[i] == '[' {
			numStack = append(numStack, curNum)
			strStack = append(strStack, ans)
			ans = ""
			curNum = 0
		} else if s[i] == ']' {
			preAns := strStack[len(strStack)-1]
			strStack = strStack[:len(strStack)-1]

			preNum := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]

			ans = preAns + strings.Repeat(ans, preNum)
		} else {
			ans += string(s[i])
		}
	}

	return ans
}
