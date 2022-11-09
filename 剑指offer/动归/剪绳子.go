func cuttingRope(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1], dp[2] = 0, 0, 1
	for i := 3; i <= n; i++ {
		for j := 1; j <= i-1; j++ {
			dp[i] = max(dp[i], max(j*(i-j), dp[i-j]*j))
		}
	}
	return dp[n]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}