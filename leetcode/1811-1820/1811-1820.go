package _1811_1820

// leetcode problem No. 1820
// https://github.com/doocs/leetcode/blob/main/solution/1800-1899/1820.Maximum%20Number%20of%20Accepted%20Invitations/README_EN.md
func maximumInvitations(grid [][]int) int {
	nBoys, nGirls := len(grid), len(grid[0])
	var visited map[int]bool     // visited[j] means whether j-th girl is explored in current round of dfs
	match := make([]int, nGirls) // match[j] means which boy the j-th girl matches
	for i := range match {
		match[i] = -1
	}

	var find func(int) bool
	find = func(i int) bool { // find function try to find a match girl for boy i
		for j, v := range grid[i] {
			if v == 1 && !visited[j] { // if can match j-th girl and the girl is not visited
				visited[j] = true
				if match[j] == -1 || find(match[j]) {
					// situation 1: match[j] == -1
					// j-th girl is not matched with a boy, so we match j-th girl with i-th boy

					// situation 2: match[j] == 1
					// j-th girl is already token by a boy(match[j]), then we try to find a new match for that boy

					match[j] = i
					return true
				}
				// note: we didn't see the visited[j] to false here
			}
		}
		return false
	}

	ans := 0
	for i := 0; i < nBoys; i++ {
		visited = map[int]bool{}
		if find(i) {
			ans++
		}
	}
	return ans
}
