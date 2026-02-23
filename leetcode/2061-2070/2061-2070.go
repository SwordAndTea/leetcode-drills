package _2061_2070

// leetcode problem No. 2062

func countVowelSubstrings(word string) int {
	vowelCharIndex := map[byte]int{
		'a': -1,
		'e': -1,
		'i': -1,
		'o': -1,
		'u': -1,
	}
	vowelSubstringStartIndex := -1 // actually the index before the start
	ans := 0
	for i := 0; i < len(word); i++ {
		if _, ok := vowelCharIndex[word[i]]; ok { // word[i] is vowel char
			vowelCharIndex[word[i]] = i
			m := i // m indicates the smallest window with all 5 vowels
			for _, index := range vowelCharIndex {
				m = min(m, index)
			}
			ans += max(0, m-vowelSubstringStartIndex)
		} else {
			vowelSubstringStartIndex = i
		}
	}
	return ans
}
