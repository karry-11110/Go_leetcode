func nthUglyNumber(n int) int {
	dp := make([]int, n)
	dp[0] = 1
	for i := 1; i < n; i++ {
		for j := 0; j <= i-1; j++ {
			if dp[j]*2 > dp[i-1] {
				dp[i] = min(dp[i], min(dp[j]*2, min(dp[j]*3, dp[j]*5)))
			} else if dp[j]*3 > dp[i-1] {
				dp[i] = min(dp[i], min(dp[j]*3, dp[j]*5))
			} else if dp[j]*5 > dp[i-1] {
				dp[i] = min(dp[i], dp[j]*5)
			}
		}
	}
	return dp[n-1]
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}