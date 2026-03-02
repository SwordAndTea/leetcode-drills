package _751_760

// leetcode problem No. 752

func openLock(deadends []string, target string) int {
	if target == "0000" {
		return 0
	}
	deadendsMap := make(map[string]bool)
	for _, deadend := range deadends {
		deadendsMap[deadend] = true
	}
	if deadendsMap[target] || deadendsMap["0000"] {
		return -1
	}
	visiteInfo := make(map[string]bool)
	q := []string{"0000"}
	visiteInfo["0000"] = true
	ans := -1
	for len(q) > 0 {
		ans++
		for i := len(q); i > 0; i-- {
			curStr := q[0]
			q = q[1:]

			if curStr == target {
				return ans
			}

			for k := 0; k < 4; k++ {
				next := make([]byte, len(curStr))
				copy(next, curStr)
				next[k] = (10+next[k]+1-'0')%10 + '0'

				if !deadendsMap[string(next)] && !visiteInfo[string(next)] {
					q = append(q, string(next))
					visiteInfo[string(next)] = true
				}

				next2 := make([]byte, len(curStr))
				copy(next2, curStr)
				next2[k] = (10+next2[k]-1-'0')%10 + '0'
				if !deadendsMap[string(next2)] && !visiteInfo[string(next2)] {
					q = append(q, string(next2))
					visiteInfo[string(next2)] = true
				}
			}
		}
	}

	return -1
}
