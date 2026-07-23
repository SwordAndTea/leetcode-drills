package _2061_2070

// leetcode problem No. 2062
func countVowelSubstrings(word string) int {
	// record the last appear index of vowel char, it serves as the right indicator of two pointer solution
	vowelCharIndex := map[rune]int{
		'a': -1,
		'e': -1,
		'i': -1,
		'o': -1,
		'u': -1,
	}

	validSubStringStartIndex := -1 // it serves as the left indicator of the two pointer solution
	ans := 0
	for i, c := range word {
		if _, ok := vowelCharIndex[c]; ok {
			vowelCharIndex[c] = i
			minIndexOfVowelChar := i // the min last appear index of vowel char
			for _, index := range vowelCharIndex {
				minIndexOfVowelChar = min(minIndexOfVowelChar, index)
			}
			ans += max(0, minIndexOfVowelChar-validSubStringStartIndex)
		} else {
			validSubStringStartIndex = i
		}
	}
	return ans
}
