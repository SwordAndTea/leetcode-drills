package _2401_2410

import "sort"

// leetcode problem No. 2410
func matchPlayersAndTrainers(players []int, trainers []int) int {
	sort.Ints(players)
	sort.Ints(trainers)
	ans := 0
	i := 0
	j := 0
	m := len(players)
	n := len(trainers)
	for i < m && j < n {
		if players[i] <= trainers[j] {
			ans++
			i++
			j++
		} else {
			j++
		}
	}
	return ans
}
