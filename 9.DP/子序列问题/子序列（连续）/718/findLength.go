package main

func findLength(A []int, B []int) int {
	result := 0
	dp := make([][]int, len(A))
	for i := 0; i < len(A); i++ {
		dp[i] = make([]int, len(B))
	}
	for i := 0; i < len(B); i++ {
		if B[i] == A[0] {
			dp[0][i] = 1
		}
	}

	for i := 0; i < len(A); i++ {
		if A[i] == B[0] {
			dp[i][0] = 1
		}
	}

	for i := 1; i < len(A); i++ {
		for j := 1; j < len(B); j++ {
			if A[i] == B[j] {
				dp[i][j] = dp[i-1][j-1] + 1
			}
			if dp[i][j] > result {
				result = dp[i][j]
			}
		}
	}

	return result
}
