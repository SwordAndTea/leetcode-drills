package dp

func LargestSequenceSum(values []int) int {
	dp := make([]int, len(values))
	dp[0] = values[0]
	max := values[0]
	for i := 1; i < len(values); i++ {
		if dp[i-1]+values[i] > values[i] {
			dp[i] = dp[i-1] + values[i]
		} else {
			dp[i] = values[i]
		}
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}

func LongestIncreasingSequence(values []int) int {
	dp := make([]int, len(values))
	dp[0] = 1
	max := 1

	for i := 1; i < len(values); i++ {
		dp[i] = 1
		for j := i - 1; j >= 0; j-- {
			if values[j] <= values[i] && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
			}
		}
		if dp[i] > max {
			max = dp[i]
		}
	}

	return max
}

func LongestCommonSubsequence(a, b string) {
	dp := make([][]int, len(a))
	for i := 0; i < len(a); i++ {
		dp[i] = make([]int, len(b))
		if a[i] == b[0] {
			dp[i][0] = 1
		} //else {
		//	dp[i][0] = 0
		//}
	}

	for j := 1; j < len(b); j++ {
		if a[0] == b[j] {
			dp[0][j] = 1
		} //else {
		//	dp[0][j] = 0
		//}
	}

	for i := 1; i < len(a); i++ {
		for j := 1; j < len(b); j++ {
			if a[i] == b[j] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				if dp[i-1][j] >= dp[i][j-1] {
					dp[i][j] = dp[i-1][j]
				} else {
					dp[i][j] = dp[i][j-1]
				}
			}
		}
	}
}

func LongestPalindromeSubsequence(a string) int {
	dp := make([][]int, len(a))
	max := 1
	for i := 0; i < len(a); i++ {
		dp[i] = make([]int, len(a))
		dp[i][i] = 1
		if i < len(a)-1 && a[i] == a[i+1] {
			dp[i][i+1] = 1
			max = 2
		} else {
			dp[i][i+1] = 0
		}
	}

	for i := 0; i < len(a); i++ {
		for j := 2; j < len(a); j++ {
			if a[i] == a[j] && dp[i+1][j-1] == 1 {
				dp[i][j] = 1
				if j-i+1 > max {
					max = j - i + 1
				}
			} // else {
			//	 dp[i][j] = 0
			// }
		}
	}

	return max
}

func OneZeroBackPack(weight []int, value []int, volume int) int {
	dp := make([][]int, len(weight))

	for i := 0; i < len(weight); i++ {
		dp[i] = make([]int, volume+1)
	}

	for i := 0; i < len(weight); i++ {
		for j := weight[i]; j <= volume; j++ {
			chose := dp[i-1][j-weight[i]] + value[i]
			notChose := dp[i-1][j]
			if chose >= notChose {
				dp[i][j] = chose
			} else {
				dp[i][j] = notChose
			}
		}
	}

	return dp[len(weight)-1][volume]
}

func CompleteBackPack(weight []int, value []int, volume int) int {
	dp := make([][]int, len(weight))

	for i := 0; i < len(weight); i++ {
		dp[i] = make([]int, volume+1)
	}

	for i := 0; i < len(weight); i++ {
		for j := weight[i]; j <= volume; j++ {
			chose := dp[i][j-weight[i]] + value[i]
			notChose := dp[i-1][j]
			if chose >= notChose {
				dp[i][j] = chose
			} else {
				dp[i][j] = notChose
			}
		}
	}

	return dp[len(weight)-1][volume]
}
