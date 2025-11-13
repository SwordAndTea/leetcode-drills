package main

func fibonacci(n int) int {
	var dp = make([]int, n+1)
	count := 0

	var solve func(i int) int
	solve = func(i int) int {
		count++
		if i == 0 || i == 1 {
			return 1
		}

		if dp[i] != 0 {
			return dp[i]
		}
		v := solve(i-1) + solve(i-2)
		dp[i] = v
		return v
	}
	res := solve(n)
	print(count)
	return res
}

func fibonacci2(n int) int {
	count := 0
	var solve func(v int) int
	solve = func(v int) int {
		count++
		if v == 0 || v == 1 {
			return 1
		}
		return solve(v-1) + solve(v-2)
	}

	res := solve(n)
	print(count)
	return res
}

func main() {

}
