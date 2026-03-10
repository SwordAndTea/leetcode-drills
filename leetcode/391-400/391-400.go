package _391_400

import (
	"slices"
	"strings"
)

// leetcode problem No. 394

func decodeString(s string) string {
	num := 0
	var strInBracket []byte

	strStack := []string{}
	numStack := []int{}

	for i := 0; i < len(s); i++ {
		c := s[i]
		if c >= '0' && c <= '9' {
			num = num*10 + int(c-'0')
		} else if c == '[' {
			numStack = append(numStack, num)
			strStack = append(strStack, string(strInBracket))
			num = 0
			strInBracket = []byte{}
		} else if c >= 'a' && c <= 'z' {
			strInBracket = append(strInBracket, c)
		} else if c == ']' {
			preNum := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]

			preStrInBracket := []byte(strStack[len(strStack)-1])
			strStack = strStack[:len(strStack)-1]

			strInBracket = slices.Concat(preStrInBracket, []byte(strings.Repeat(string(strInBracket), preNum)))
		}
	}
	return string(strInBracket)
}
