// 1.定义：dp[j]表示容量（这里说容量更形象，其实就是重量）为j的背包，最多可以背dp[j]这么重的石头。
// 2.递推式：dp[j] = max(dp[j], dp[j - weight[i]] + value[i])
// 3.初始理论： dp[j]中的j表示容量，那么最大容量（重量）是多少呢，就是所有石头的重量和。
// 				因为重量都不会是负数，所以dp[j]都初始化为0就可以了，这样在递归公式dp[j] = max(dp[j], dp[j - stones[i]] + stones[i]);中dp[j]才不会初始值所覆盖。
// 4.外层遍历物品，内层倒序遍历背包
// 5.推导
package main

func lastStoneWeightII(stones []int) int {
	//先求容量target
	sum := 0
	for _, value := range stones {
		sum += value
	}
	target := sum / 2
	//定义
	dp := make([]int, 15001) //15001 = 30*1000/2 + 1
	//递推
	for i := 0; i < len(stones); i++ {
		for j := target; j >= stones[i]; j-- {
			dp[j] = max(dp[j], dp[j-stones[i]]+stones[i])
		}
	}
	return sum - 2*dp[target]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
