//dp[i][j]代表i个骰子点数和为j的方案数
func dicesProbability(n int) []float64 {
	dp := make([][]int, n+1)
	for i := range dp {
		if i != 0 {
			dp[i] = make([]int, 6*n+1)
		}
	}
	for i := 1; i <= 6; i++ {
		dp[1][i] = 1
	}
	for i := 2; i <= n; i++ {
		for j := i; j <= 6*i; j++ {
			for k := 1; k <= 6 && k < j; k++ {
				dp[i][j] += dp[i-1][j-k]
			}
		}
	}
	result := make([]float64, 0)
	sum := 0
	for i := n; i <= 6*n; i++ {
		sum += dp[n][i]
	}

	for i := n; i <= 6*n; i++ {
		result = append(result, float64(dp[n][i])/float64(sum))
	}
	return result
}