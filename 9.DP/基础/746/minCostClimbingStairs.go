// 1.定义：dp[i]： 爬到第i层楼梯的最小花费
// 2.递推式：dp[i] = min(dp[i-2], dp[i-1]) + cost[i]
// 3.初始化：dp[0]=cost[0],dp[1]=cost[1]
// 4.遍历顺序：从前往后
// 5.举例推导 ，debug就是打印dp数组
func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost))
	dp[0], dp[1] = cost[0], cost[1]
	for i := 2; i <= len(cost)-1; i++ {
		dp[i] = min(dp[i-2], dp[i-1]) + cost[i]
	}
	// 注意最后一步可以理解为不用花费，所以取倒数第一步，第二步的最少值
	return min(dp[len(cost)-1], dp[len(cost)-2])
}
func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}