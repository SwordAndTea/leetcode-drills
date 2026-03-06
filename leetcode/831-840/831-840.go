package _831_840

// leetcode problem No. 834

func sumOfDistancesInTree(n int, edges [][]int) []int {
	connectMap := make(map[int][]int)
	distanceSum := make([]int, n)
	childrenCount := make([]int, n)
	for _, edge := range edges {
		connectMap[edge[0]] = append(connectMap[edge[0]], edge[1])
		connectMap[edge[1]] = append(connectMap[edge[1]], edge[0])
	}

	// dfs1 counts the number of total children under node and the total distance to its children for the node
	var dfs1 func(node, parent int)
	dfs1 = func(node, parent int) {
		for _, child := range connectMap[node] {
			if child == parent {
				continue
			}
			dfs1(child, node)
			childrenCount[node] += childrenCount[child]
			distanceSum[node] += distanceSum[child] + childrenCount[child]
		}
		childrenCount[node]++ // add self
	}

	var dfs2 func(root, parent int)
	dfs2 = func(root, parent int) {
		for _, child := range connectMap[root] {
			if child == parent {
				continue
			}
			// set the children as new root
			distanceSum[child] = distanceSum[root] - childrenCount[child] + n - childrenCount[child]
			dfs2(child, root)
		}
	}

	dfs1(0, -1)
	dfs2(0, -1)

	return distanceSum
}
