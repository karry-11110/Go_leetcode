// 1.含义：dp[i]： 1到i为节点组成的二叉搜索树的个数为dp[i]。
// 2.递推式：dp[i] += dp[j - 1] * dp[i - j];
// j-1 为j为头结点左子树节点数量，i-j 为以j为头结点右子树节点数量
// 3.初始化：dp[0] = 1
// 4.遍历顺序：首先一定是遍历节点数,遍历i里面每一个数作为头结点的状态，用j来遍历。
// 5.推导dp
package main

func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}
