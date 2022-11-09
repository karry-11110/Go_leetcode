package main

import "fmt"
func main() {
	row, col := 0, 0
	fmt.Scanln(&row, &col)
	grid := make([][]int, row)
	for index,_ := range grid{
		grid[index] = make([]int, col)
		for i := 0; i < col; i++ {
			m := 0
			fmt.Scan(&m)
			grid[index][i] = m
		}
	}
	fmt.Println(maxValue(grid))
}
func maxValue(grid [][]int) int {
	dp := make([][]int, len(grid))
	for i, _ := range dp {
		dp[i] = make([]int, len(grid[0]))
	}
	dp[0][0] = grid[0][0]
	for i := 1; i < len(grid[0]); i++ {
		dp[0][i] = dp[0][i-1] + grid[0][i]
	}
	for i := 1; i < len(grid); i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[0]); j++ {
			dp[i][j] = max(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[len(grid)-1][len(grid[0])-1]
}
func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}