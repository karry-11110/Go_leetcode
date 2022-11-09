// 1. 在一维dp数组中，dp[j]表示：容量为j的背包，所背的物品价值可以最大为dp[j]
// 2. dp[j] = max(dp[j], dp[j - weight[i]] + value[i]);
// dp[j]有两个选择，一个是取自己dp[j] 相当于 二维dp数组中的dp[i-1][j]，即不放物品i，
// 				一个是取dp[j - weight[i]] + value[i]，即放物品i，指定是取最大的，毕竟是求最大价值，
// 3.和dp数组的定义吻合：dp数组在推导的时候一定是取价值最大的数，如果题目给的价值都是正整数那么非0下标都初始化为0就可以了。
// 这样才能让dp数组在递归公式的过程中取的最大的价值，而不是被初始值覆盖了。假设物品价值都是大于0的，所以dp数组初始化的时候，都初始为0就可以了。
// 4.必须先遍历物品再倒序遍历背包。
// 5.推导
package main

import "fmt"

func test_1_bagProblem(weight, value []int, bagWeight int) int {
	//定义并且初始化
	dp := make([]int, bagWeight+1)
	for i := 0; i < len(weight); i++ {
		// 这里必须倒序,区别二维,因为二维dp保存了i的状态
		for j := bagWeight; j >= weight[i]; j-- {
			// 递推公式
			dp[j] = max(dp[j], dp[j-weight[i]]+value[i])
		}
	}
	return dp[bagWeight]
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
	fmt.Println(test_1_bagProblem(weight, value, 4))
}
