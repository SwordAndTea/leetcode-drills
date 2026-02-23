package _1831_1840

// leetcode problem No. 1839

func longestBeautifulSubstring(word string) int {
	ans := 0
	i := 0
	for i < len(word) {
		if word[i] == 'a' {
			preChar := byte('a')
			j := i + 1
			for j < len(word) {
				shouldBreak := false
				switch word[j] {
				case 'a':
					if preChar != 'a' {
						shouldBreak = true
					}
				case 'e':
					if preChar != 'e' && preChar != 'a' {
						shouldBreak = true
					}
				case 'i':
					if preChar != 'i' && preChar != 'e' {
						shouldBreak = true
					}
				case 'o':
					if preChar != 'o' && preChar != 'i' {
						shouldBreak = true
					}
				case 'u':
					if preChar != 'u' && preChar != 'o' {
						shouldBreak = true
					} else {
						ans = max(ans, j-i+1)
					}
				}
				if shouldBreak {
					break
				}
				preChar = word[j]
				j++
			}
			i = j
		} else {
			i++
		}
	}
	return ans
}
