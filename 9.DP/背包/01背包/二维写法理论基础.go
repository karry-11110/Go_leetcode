// 对于背包问题
// 1. dp[i][j]：表示从下标为[0-i]的物品里任意取，放进容量为j的背包，价值总和最大是多少
// 2. dp[i][j] = max(dp[i - 1][j], dp[i - 1][j - weight[i]] + value[i]);
// 3.和dp数组的定义吻合：背包容量j为0的话，即dp[i][0]，无论是选取哪些物品，背包价值总和一定为0。
//                     j < weight[0]的时候，dp[0][j] 应该是 0，j >= weight[0]时，dp[0][j] 应该是value[0]
// 4.先遍历 物品还是先遍历背包重量呢？其实都可以（递归的本质和递推的方向）！！ 但是先遍历物品更好理解
// 5.推导
package main

import "fmt"

func test_2_bagProblem(weight, value []int, bagweight int) int {
	//定义dp数组
	dp := make([][]int, len(weight))
	for i := range dp {
		dp[i] = make([]int, bagweight+1)
	}
	//初始化
	for j := 0; j <= bagweight; j++ {
		if j >= weight[0] {
			dp[0][j] = value[0]
		}
	}
	//递推式

	for i := 1; i < len(weight); i++ {
		for j := 0; j < bagweight+1; j++ {
			if j < weight[i] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-weight[i]]+value[i])
			}
		}
	}
	return dp[len(weight)-1][bagweight]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {
	weight := []int{1, 3, 4}
	value := []int{15, 20, 30}
	fmt.Println(test_2_bagProblem(weight, value, 4))
}
