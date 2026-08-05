package _3301_3310

// leetcode problem No. 3310
func remainingMethods(n int, k int, invocations [][]int) []int {
	callee := make(map[int][]int)
	caller := make(map[int][]int)
	removeGroup := make(map[int]bool)

	for _, inv := range invocations {
		callee[inv[0]] = append(callee[inv[0]], inv[1])
		caller[inv[1]] = append(caller[inv[1]], inv[0])
	}

	visited := make(map[int]bool)

	bfs := func(start int) {
		queue := make([]int, 0, n)
		queue = append(queue, start)
		for len(queue) > 0 {
			cur := queue[0]
			queue = queue[1:]
			removeGroup[cur] = true
			for _, aCallee := range callee[cur] {
				if !visited[aCallee] {
					queue = append(queue, aCallee)
					visited[aCallee] = true
				}
			}
		}
	}

	bfs(k)

	shouldRemove := true
	for i := 0; i < n; i++ {
		if removeGroup[i] {
			for _, aCaller := range caller[i] {
				if !removeGroup[aCaller] {
					shouldRemove = false
					break
				}
			}
		}

		if !shouldRemove {
			break
		}
	}
	ans := make([]int, 0, n)
	for i := 0; i < n; i++ {
		if shouldRemove && removeGroup[i] {
			continue
		}
		ans = append(ans, i)
	}

	return ans
}
