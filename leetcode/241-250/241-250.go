package _241_250

// leetcode problem No. 249
// https://github.com/doocs/leetcode/blob/main/solution/0200-0299/0249.Group%20Shifted%20Strings/README_EN.md
func groupStrings(strings []string) [][]string {
	strHash := func(s string) string { // the hash function will convert s to start with 'a'
		// I think this hush function is a great candidate for rotate loop, for example from 0 to 9 then 9 -> 0
		ans := []byte(s)
		diff := ans[0] - 'a'
		for i := 0; i < len(s); i++ {
			ans[i] -= diff
			if ans[i] < 'a' {
				ans[i] += 26
			}
		}
		return string(ans)
	}

	groupMap := make(map[string][]string)
	for _, s := range strings {
		sHash := strHash(s)
		groupMap[sHash] = append(groupMap[sHash], s)
	}

	ans := [][]string{}

	for _, group := range groupMap {
		ans = append(ans, group)
	}

	return ans
}
