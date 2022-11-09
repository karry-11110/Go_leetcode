package main

// dp[j]：凑成总金额j的货币组合数为dp[j]
// dp[j] （考虑coins[i]的组合总和） 就是所有的dp[j - coins[i]]（不考虑coins[i]）相加:dp[j] += dp[j - coins[i]];
// dp[0]一定要为1，dp[0] = 1,下标非0的dp[j]初始化为0
// 求组合数：先物品再背包
func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			dp[j] += dp[j-coins[i]]
		}
	}
	return dp[amount]
}
