package _1411_1420

// leetcode problem No. 1411
func numOfWays(n int) int {
	modulo := 1_000_000_000 + 7
	same := 6
	diff := 6
	sameToSame := 3 // the 1x3 block that contains same color can lead to 3 types of block that contains same color in next level
	sameToDiff := 2 // the 1x3 block that contains same color can lead to 2 types of block that contains different color in next level
	diffToSame := 2 // the 1x3 block that contains different color can lead to 2 types of block that contains different color in next level
	diffToDiff := 2 // the 1x3 block that contains different color can lead to 2 types of block that contains same color in next level
	for i := 2; i < n; i++ {
		nextSame := (same*sameToSame + diff*diffToSame) % modulo
		nextDiff := (same*sameToDiff + diff*diffToDiff) % modulo
		same = nextSame
		diff = nextDiff
	}
	return same + diff
}
