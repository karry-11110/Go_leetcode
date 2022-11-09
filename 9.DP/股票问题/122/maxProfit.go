package main

// //贪心算法
// func maxProfit(prices []int) int {
//     var sum int
//     for i := 1; i < len(prices); i++ {
//         // 累加每次大于0的交易
//         if prices[i]-prices[i-1] > 0 {
//             sum += prices[i]-prices[i-1]
//         }
//     }
//     return sum
// }
func maxProfit(prices []int) int {
	//创建数组
	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = -prices[0]
	dp[0][1] = 0
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
	}
	return dp[len(prices)-1][1]
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
