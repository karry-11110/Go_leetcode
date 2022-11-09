// 1.定义：dp[i]的定义为：第i个数的斐波那契数值为dp[i]
// 2.递推式：dp[i] = dp[i - 1] + dp[i - 2];
// 3.初始化：dp[0] = 0;dp[1] = 1;
// 4.遍历顺序：从前往后
// 5.举例推导 ，debug就是打印dp数组

func fib(n int) int {
	if n < 2 {
		return n
	}
	dp := make([]int, n+1) //只能定义切片，并且要初始化长度
	dp[0], dp[1] = 0, 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
