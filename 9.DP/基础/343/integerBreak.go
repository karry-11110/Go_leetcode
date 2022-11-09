// 1.含义：dp[i]：分拆数字i，可以得到的最大乘积为dp[i]。
// 2.递推式：dp[i] = max(dp[i], max((i - j) * j, dp[i - j] * j));
// j * (i - j) 是单纯的把整数拆分为两个数相乘，而j * dp[i - j]是拆分成两个以及两个以上的个数相乘。
// 3.初始化：dp[2] = 1
// 4.遍历顺序：从前往后
// 5.推导
package main

func integerBreak(n int) int {
	dp := make([]int, n+1)
	dp[2] = 1
	for i := 3; i < n+1; i++ {
		for j := 1; j <= i-1; j++ { //这里也可以不写=，因为dp[1]=0
			dp[i] = max(dp[i], max(j*(i-j), j*dp[i-j])) //这种写法就是判断一轮循环中，总取dp[i]最大值
		}
	}
	return dp[n]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
