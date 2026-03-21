package _801_810

// leetcode problem No. 804

func uniqueMorseRepresentations(words []string) int {
	morseArray := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

	morseCode := func(w string) string {
		result := ""
		for i := 0; i < len(w); i++ {
			result += morseArray[w[i]-'a']
		}
		return result
	}

	transformMap := map[string]bool{}
	for _, word := range words {
		transformMap[morseCode(word)] = true
	}

	return len(transformMap)
}
