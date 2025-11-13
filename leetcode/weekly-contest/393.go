package weekly_contest

func findLatestTime(s string) string {
	tmp := []byte(s)
	pre, post := tmp[0:2], tmp[3:5]
	if pre[0] == '?' && pre[1] == '?' {
		pre[0] = '1'
		pre[1] = '1'
	} else if pre[0] == '?' {
		if pre[1] > '1' {
			pre[0] = '0'
		} else {
			pre[0] = '1'
		}
	} else if pre[1] == '?' {
		if pre[0] == '0' {
			pre[1] = '9'
		} else {
			pre[1] = '1'
		}
	}

	if post[0] == '?' && post[1] == '?' {
		post[0] = '5'
		post[1] = '9'
	} else if post[0] == '?' {
		post[0] = '5'
	} else if post[1] == '?' {
		post[1] = '9'
	}

	return string(tmp)
}

var primeList = map[int]bool{
	2:  true,
	3:  true,
	5:  true,
	7:  true,
	11: true,
	13: true,
	17: true,
	19: true,
	23: true,
	29: true,
	31: true,
	37: true,
	41: true,
	43: true,
	47: true,
	53: true,
	59: true,
	61: true,
	67: true,
	71: true,
	73: true,
	79: true,
	83: true,
	89: true,
	97: true,
}

func maximumPrimeDifference(nums []int) int {
	prime1, prime2 := -1, -1

	for i := 0; i < len(nums); i++ {
		if primeList[nums[i]] {
			if prime1 == -1 {
				prime1 = i
			} else {
				prime2 = i
			}
		}
	}

	if prime2 <= prime1 {
		return 0
	}

	return prime2 - prime1
}

func GCD(a, b int) int {
	if b == 0 {
		return a
	}

	return GCD(b, a%b)
}

func LCM(a, b int) int {
	gcd := GCD(a, b)

	return a / gcd * b
}

// bitOneCount how many 1 in the binary from of x, x > 0
func bitOneCount(x int) int {
	result := 0
	for x != 0 {
		x &= x - 1
		result++
	}
	return result
}

func findKthSmallest(coins []int, k int) int64 {
	n := len(coins)
	nn := 1 << n
	count := func(x int64) int64 {
		// count of values less than or equal to x which are divisible by at least one number from the coins
		// it is actually the ordered index of x by using those coins
		cnt := int64(0)
		for i := 1; i < nn; i++ { // for every coin usage situation
			lcm := 1
			for j := 0; j < n; j++ {
				if i&(1<<j) > 0 { // coin[j] used
					lcm = LCM(lcm, coins[j])
				}
			}
			if bitOneCount(i)%2 == 0 { // used even number of coins, see more in Inclusion-Exclusion Principle
				cnt -= x / int64(lcm)
			} else {
				cnt += x / int64(lcm)
			}
		}
		return cnt
	}

	min := coins[0]
	for i := 1; i < n; i++ {
		if coins[i] < min {
			min = coins[i]
		}
	}

	result := int64(0)
	left, right := int64(min), int64(min)*int64(k)
	for left <= right {
		mid := (left + right) / 2
		index := count(mid) // get the index of mid
		// if index == k, mid is not necessarily the target value as it may not constructable by those coins
		if index < int64(k) {
			left = mid + 1
		} else {
			result = mid
			right = mid - 1
		}
	}
	return result
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func minimumValueSum(nums []int, andValues []int) int {
	n, m := len(nums), len(andValues)

	var solve func(curI, curJ, mask int) int
	memo := make(map[int]map[int]map[int]int)

	solve = func(curI, curJ, mask int) int {
		if curI == n && curJ == m {
			return 0
		}
		if curI == n || curJ == m {
			return 1e8
		}

		subMemo := memo[curI]
		if subMemo == nil {
			subMemo = make(map[int]map[int]int)
		}
		memo[curI] = subMemo

		nextMemo := subMemo[curJ]
		if nextMemo == nil {
			nextMemo = make(map[int]int)
		}
		subMemo[curJ] = nextMemo

		if val, ok := nextMemo[mask]; ok {
			return val
		}

		mask &= nums[curI]
		if mask < andValues[curJ] {
			return 1e8
		}

		if mask == andValues[curJ] {
			res := min(solve(curI+1, curJ, mask), nums[curI]+solve(curI+1, curJ+1, -1))
			memo[curI][curJ][mask] = res
			return res
		}

		res := solve(curI+1, curJ, mask)
		memo[curI][curJ][mask] = res
		return res
	}

	result := solve(0, 0, -1) // -1's binary format is all 1 in signed int
	if result == 1e8 {
		return -1
	}
	return result
}
