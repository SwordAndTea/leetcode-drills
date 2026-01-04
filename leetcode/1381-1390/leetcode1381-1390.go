package _1381_1390

import "math"

// leetcode problem No. 1390
func sumForDivisor(nums []int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		curNum := nums[i]
		numberOfDivisor := 1
		curSum := 1
		if curNum > 1 {
			numberOfDivisor = 2
			curSum += curNum
		}

		for j := 2; j <= int(math.Sqrt(float64(curNum))); j++ {
			if curNum%j == 0 {
				k := curNum / j
				if k == j {
					numberOfDivisor += 1
					curNum += j
				} else {
					numberOfDivisor += 2
					curSum += j + k
				}
			}

			if numberOfDivisor > 4 {
				break
			}
		}

		if numberOfDivisor == 4 {
			sum += curSum
		}
	}
	return sum
}
