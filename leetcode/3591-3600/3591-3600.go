package _3591_3600

import "slices"

// leetcode problem No. 3593
func minIncrease(n int, edges [][]int, cost []int) int {
	tree := make(map[int][]int, n+1)
	for _, edge := range edges {
		tree[edge[0]] = append(tree[edge[0]], edge[1])
		tree[edge[1]] = append(tree[edge[1]], edge[0])
	}
	ans := 0
	var dfs func(curNode int, parent int) int // dfs get the cost from curNode to all it's children
	dfs = func(curNode int, parent int) int {
		costByLevel := []int{}
		for _, child := range tree[curNode] {
			if child == parent {
				continue
			}
			costByLevel = append(costByLevel, dfs(child, curNode))
		}
		if len(costByLevel) == 0 {
			return cost[curNode]
		}
		if len(costByLevel) == 1 {
			return costByLevel[0] + cost[curNode]
		}
		maxCostValue := slices.Max(costByLevel)
		for _, v := range costByLevel {
			if v != maxCostValue {
				ans++
			}
		}
		return maxCostValue + cost[curNode] //note: we did not multiple maxCostValue by number of node
	}

	dfs(0, -1)
	return ans
}
