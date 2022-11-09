package main

import "math"

func numSquares(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt32
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if i >= j*j {
				dp[i] = min(dp[i], dp[i-j*j]+1)
			}
		}
	}

	return dp[n]
}
func min(i, j int) int {
	if i < j {
		return i
	}

	return j
}
