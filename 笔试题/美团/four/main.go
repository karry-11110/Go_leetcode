package main
import "fmt"
func main() {
	n, m := 0, 0
	fmt.Scanln(&n, &m)
	A, B := make([]int, n), make([]int, m)
	for i := 0; i < n; i++ {
		fmt.Scan(&A[i])
	}
	for i := 0; i < m; i++ {
		fmt.Scan(&B[i])
	}
	res := solve(n, m, A, B)
	fmt.Println(res)
}
func solve(n, m int, A, B []int) int {
	dp := make([][]int, n+1)
	for i , _ := range dp {
		dp[i] = make([]int, m+1)
	}
	for i := 1; i < n+1; i++ {
		dp[i][0] = dp[i-1][0] + abs(A[i-1])
	}
	for j := 1; j < m+1; j++ {
		dp[0][j] = dp[0][j-1] + abs(B[j-1])
	}
	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			dp[i][j] = min(min(dp[i-1][j]+abs(A[i-1]), dp[i][j-1]+abs(B[j-1])), dp[i-1][j-1]+abs(A[i-1]-B[j-1]))
		}
	}
	return dp[n][m]
}
func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
func min(i, j int) int {
	if i < j {
		return i
	}
	return  j
}
