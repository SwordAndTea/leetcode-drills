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
		costOfChildren := []int{}
		for _, child := range tree[curNode] {
			if child == parent {
				continue
			}
			costOfChildren = append(costOfChildren, dfs(child, curNode))
		}
		if len(costOfChildren) == 0 {
			return cost[curNode]
		}
		if len(costOfChildren) == 1 {
			return costOfChildren[0] + cost[curNode]
		}
		maxCostValue := slices.Max(costOfChildren)
		for _, v := range costOfChildren {
			if v != maxCostValue {
				ans++
			}
		}
		return maxCostValue + cost[curNode] //note: we did not multiple maxCostValue by number of node
	}

	dfs(0, -1)
	return ans
}
