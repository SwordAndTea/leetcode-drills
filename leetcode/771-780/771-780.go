package _771_780

// leetcode problem No. 777
func canTransform(start string, result string) bool {
	m, n := len(start), len(result)
	i, j := 0, 0
	for i < m && j < n {
		for i < m && start[i] == 'X' { // find next 'L' or 'R' in start string
			i++
		}
		for j < n && result[j] == 'X' { // find next 'L' or 'R' in result string
			j++
		}
		if i >= m && j >= n {
			break
		}

		if i < m && j >= n {
			return false
		}

		if i >= m && j < n {
			return false
		}

		if start[i] != result[j] {
			return false
		}

		if start[i] == 'L' && i < j {
			return false
		}

		if start[i] == 'R' && i > j {
			return false
		}
		i++
		j++
	}

	for i < m {
		if start[i] != 'X' {
			return false
		}
		i++
	}

	for j < n {
		if result[j] != 'X' {
			return false
		}
		j++
	}

	return true
}
