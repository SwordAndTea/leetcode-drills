package dp

func LargestSequenceSum(values []int) int {
	// dpArray[i] stands for the largest sequences sum that ending with values[i]
	dpArray := make([]int, len(values))
	dpArray[0] = values[0]
	maxValue := values[0]
	for i := 1; i < len(values); i++ {
		if dpArray[i-1] > 0 {
			dpArray[i] = dpArray[i-1] + values[i]
		} else {
			dpArray[i] = values[i]
		}
		if dpArray[i] > maxValue {
			maxValue = dpArray[i]
		}
	}
	return maxValue
}

func LongestIncreasingSequence(values []int) int {
	// dpArray[i] stands for the length of longest increasing sequence that ends withs values[i]
	dpArray := make([]int, len(values))
	dpArray[0] = 1
	maxLength := 1
	for i := 1; i < len(values); i++ {
		dpArray[i] = 1
		for j := i - 1; j >= 0; j-- {
			if values[i] > values[j] {
				dpArray[i] = max(dpArray[i], dpArray[j]+1)
			}
		}
		if dpArray[i] > maxLength {
			maxLength = dpArray[i]
		}
	}
	return maxLength
}

func LongestCommonSubsequence(str1, str2 string) int {
	// dpArray[i][j] stands for the length of the longest common subsequence of str1[0: i) and str[0: j)
	dpArray := make([][]int, len(str1)+1)
	for i := range dpArray {
		dpArray[i] = make([]int, len(str2)+1)
	}

	for i := 1; i <= len(str1); i++ {
		for j := 1; j <= len(str2); j++ {
			if str1[i-1] == str2[j-1] {
				dpArray[i][j] = dpArray[i-1][j-1] + 1
			} else {
				dpArray[i][j] = max(dpArray[i-1][j], dpArray[i][j-1])
			}
		}
	}

	return dpArray[len(str1)][len(str2)]
}

func LongestPalindromeSubsequence(a string) int {
	// dpArray[i][j] stands for the length of the palindrome of a[i: j] (right boundary inclusive)
	dpArray := make([][]int, len(a))
	maxValue := 1
	for i := range dpArray {
		dpArray[i] = make([]int, len(a))
		dpArray[i][i] = 1
		if i+1 < len(a) && a[i+1] == a[i] {
			dpArray[i][i+1] = 2
			maxValue = 2
		}
	}
	// NOTE: if when we are calculating dpArray[0][4] and a[0] == a[4]
	// 	the dpArray[0][4] will be derivative from dpArray[1][3], which is not known yet
	//for i := 0; i <= len(a); i++ {
	//	for j := i + 2; j <= len(a); j++ {
	//		if a[i] == a[j] {
	//			dpArray[i][j] = dpArray[i+1][j-1] + 2
	//		}
	//		if dpArray[i][j] > maxValue {
	//			maxValue = dpArray[i][j]
	//		}
	//	}
	//}

	// We do calculation by starting solve the string length of 3, then 4, 5, etc.
	for strLen := 3; strLen < len(a); strLen++ { // the length of string
		for i := 0; i+strLen-1 < len(a); i++ {
			j := i + strLen - 1
			if a[i] == a[j] && dpArray[i+1][j-1] != 0 {
				dpArray[i][j] = dpArray[i+1][j-1] + 2
				if dpArray[i][j] > maxValue {
					maxValue = dpArray[i][j]
				}
			}
		}
	}
	return maxValue
}

func OneZeroBackPack(weight []int, value []int, volume int) int {
	numberOfItem := len(weight)
	// dpArray[i][j] stand for the must value of putting first i number of items into a pack of volume j
	dpArray := make([][]int, numberOfItem)
	for i := range dpArray {
		dpArray[i] = make([]int, volume+1)
	}

	for i := 0; i <= numberOfItem; i++ {
		for j := weight[i]; j <= volume; j++ {
			dpArray[i][j] = max(dpArray[i-1][j] /*not choose item i*/, value[i]+dpArray[i-1][j-weight[i]] /*choose item i*/)
		}
	}
	return dpArray[numberOfItem-1][volume]
}

func CompleteBackPack(weight []int, value []int, volume int) int {
	numberOfItem := len(weight)
	// dpArray[i][j] stand for the must value of putting first i number of items into a pack of volume j

	dpArray := make([][]int, numberOfItem)
	for i := range dpArray {
		dpArray[i] = make([]int, volume+1)
	}

	for i := 0; i <= numberOfItem; i++ {
		for j := weight[i]; j <= volume; j++ {
			dpArray[i][j] = max(dpArray[i-1][j] /*not choose item i*/, value[i]+dpArray[i][j-weight[i]] /*choose item i*/)
		}
	}
	return dpArray[numberOfItem-1][volume]
}
