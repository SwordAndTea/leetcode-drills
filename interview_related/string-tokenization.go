package interview_related

import "strings"

func stringTokenization(text string, dictionary []string) []string {
	i := 0
	var ans []string
	n := len(text)
	dictionaryMap := make(map[string]string)
	for _, d := range dictionary {
		dList := strings.Split(d, ":")
		dictionaryMap[dList[0]] = dList[1]
	}
	for i < n {
		maxMatchLen := 0
		maxMatchKey := ""
		for k, _ := range dictionaryMap {
			keyLen := len(k)
			if i+keyLen <= n && text[i:i+keyLen] == k {
				if keyLen > maxMatchLen {
					maxMatchLen = keyLen
					maxMatchKey = k
				}
			}
		}
		if maxMatchKey != "" {
			ans = append(ans, dictionaryMap[maxMatchKey])
			i += len(maxMatchKey)
		} else {
			ans = append(ans, text[i:i+1])
			i++
		}
	}
	return ans
}
