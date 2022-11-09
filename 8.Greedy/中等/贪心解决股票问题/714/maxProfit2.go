package main

// 在做收获利润操作的时候其实有三种情况：

// 情况一：收获利润的这一天并不是收获利润区间里的最后一天（不是真正的卖出，相当于持有股票），所以后面要继续收获利润。
// 情况二：前一天是收获利润区间里的最后一天（相当于真正的卖出了），今天要重新记录最小价格了。
// 情况三：不作操作，保持原有状态（买入，卖出，不买不卖）
func maxProfit(prices []int, fee int) int {
	result := 0
	minPrice := prices[0] //记录最低价格，即为当前利润加本金，当你的价格高于最低加fee，才可能收获新的利润
	for i := 1; i < len(prices); i++ {
		//情况2
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
		//情况3

		//计算利润
		if prices[i] > minPrice+fee {
			result += prices[i] - minPrice - fee
			minPrice = prices[i] - fee // 情况一，这一步很关键!!!!!!!经典
		}
	}
	return result

}
