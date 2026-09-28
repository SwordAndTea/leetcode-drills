package _1611_1620

// leetcode problem No. 1614
func maxDepth(s string) int {
	var stack []int
	ans := 0
	for _, c := range s {
		if c == '(' {
			stack = append(stack, 0)
			if len(stack) > ans {
				ans = len(stack)
			}
		} else if c == ')' {
			stack = stack[:len(stack)-1]
		}
	}
	return ans
}
