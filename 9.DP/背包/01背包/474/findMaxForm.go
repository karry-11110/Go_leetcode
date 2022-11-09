// 这就是一个典型的01背包！ 只不过物品的重量有了两个维度而已
// dp[i][j]：最多有i个0和j个1的strs的最大子集的大小为dp[i][j]。
// dp[i][j] = max(dp[i][j], dp[i - zeroNum][j - oneNum] + 1);
// 因为物品价值不会是负数，初始为0，保证递推的时候dp[i][j]不会被初始值覆盖
// 外层for循环遍历物品，内层for循环遍历背包容量且从后向前遍历！:物品就是strs里的字符串，背包容量就是题目描述中的m和n
package main

func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i < len(strs); i++ {
		zeroNum, oneNum := 0, 0
		//计算0,1 个数
		//或者直接strings.Count(strs[i],"0")
		for _, v := range strs[i] {
			if v == '0' {
				zeroNum++
			}
		}
		oneNum = len(strs[i]) - zeroNum
		// 从后往前 遍历背包容量
		for j := m; j >= zeroNum; j-- {
			for k := n; k >= oneNum; k-- {
				// 推导公式
				dp[j][k] = max(dp[j][k], dp[j-zeroNum][k-oneNum]+1)
			}
		}

	}
	return dp[m][n]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
