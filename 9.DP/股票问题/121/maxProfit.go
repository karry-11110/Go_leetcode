package main

// //贪心
// func maxProfit(prices []int) int {
// 	low := math.MaxInt32
// 	result := 0
// 	for i := 0; i < len(prices); i++ {
// 		low = min(low, prices[i])
// 		result = max(result, prices[i]-low)
// 	}
// 	return result
// }
// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }
// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }

//动态规划
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = -prices[0]
	dp[0][1] = 0
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], -prices[i])
		dp[i][1] = max(dp[i-1][0]+prices[i], dp[i-1][1])
	}
	return dp[len(prices)-1][1]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
