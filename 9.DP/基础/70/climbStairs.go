// 1.定义：dp[i]： 爬到第i层楼梯，有dp[i]种方法
// 2.递推式：dp[i] = dp[i - 1] + dp[i - 2];
// 3.初始化：不考虑dp[0]如果初始化，只初始化dp[1] = 1，dp[2] = 2，然后从i = 3开始递推，这样才符合dp[i]的定义。
// 4.遍历顺序：从前往后
// 5.举例推导 ，debug就是打印dp数组

func climbStairs(n int) int {

	if n < 2 {
		return n
	}
	dp := make([]int, n) //只能定义切片，并且要初始化长度
	dp[1], dp[2] = 1, 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]

}