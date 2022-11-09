package main

//贪心
// 在做收获利润操作的时候其实有三种情况：

// 情况一：收获利润的这一天并不是收获利润区间里的最后一天（不是真正的卖出，相当于持有股票），所以后面要继续收获利润。
// 情况二：前一天是收获利润区间里的最后一天（相当于真正的卖出了），今天要重新记录最小价格了。
// 情况三：不作操作，保持原有状态（买入，卖出，不买不卖）
// func maxProfit(prices []int, fee int) int {
// 	result := 0
// 	minPrice := prices[0] //记录最低价格，即为当前利润加本金，当你的价格高于最低加fee，才可能收获新的利润
// 	for i := 1; i < len(prices); i++ {
// 		//情况2
// 		if prices[i] < minPrice {
// 			minPrice = prices[i]
// 		}
// 		//情况3
// 		//计算利润
// 		if prices[i] > minPrice+fee {
// 			result += prices[i] - minPrice - fee
// 			minPrice = prices[i] - fee // 情况一，这一步很关键!!!!!!!经典
// 		}
// 	}
// 	return result
// }

//动态规划
func maxProfit(prices []int, fee int) int {
	//创建数组
	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = -prices[0]
	dp[0][1] = 0
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i]-fee)
	}
	return dp[len(prices)-1][1]
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
