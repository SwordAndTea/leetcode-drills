package _1921_1930

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// leetcode problem No. 1927

func sumGame(num string) bool {
	n := len(num)
	leftSum := 0
	leftQuestionMark := 0
	rightSum := 0
	rightQuestionMark := 0

	for i := 0; i < n/2; i++ {
		if num[i] == '?' {
			leftQuestionMark++
		} else {
			leftSum += int(num[i] - '0')
		}
	}
	for i := n / 2; i < n; i++ {
		if num[i] == '?' {
			rightQuestionMark++
		} else {
			rightSum += int(num[i] - '0')
		}
	}

	if (leftQuestionMark+rightQuestionMark)%2 == 1 {
		return true
	}

	if leftQuestionMark == rightQuestionMark {
		return leftSum != rightSum
	}

	roundOfPlay := (leftQuestionMark - rightQuestionMark) / 2 // no equal 0
	diff := leftSum - rightSum
	if roundOfPlay*diff >= 0 {
		return true
	}
	return abs(diff) != abs(roundOfPlay*9)
}
