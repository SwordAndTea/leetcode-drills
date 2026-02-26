package _311_320

// leetcode problem No. 316

func removeDuplicateLetters(s string) string {
	var stack []rune
	inStack := map[rune]bool{}
	lastIndex := map[rune]int{}

	for i, c := range s {
		lastIndex[c] = i
	}

	for i, c := range s {
		if !inStack[c] {
			// if the top char in the stack is greater than current char and it will appear after
			// pop it from the stack
			for len(stack) > 0 && c < stack[len(stack)-1] && i < lastIndex[stack[len(stack)-1]] {
				inStack[stack[len(stack)-1]] = false
				stack = stack[:len(stack)-1] // pop
			}
			inStack[c] = true
			stack = append(stack, c) // push
		}
	}
	return string(stack)
}
