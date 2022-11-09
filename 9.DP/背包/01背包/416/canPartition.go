// 1.定义：dp[j]表示 背包总容量是j，最大可以凑成j的子集总和为dp[j]
// 2.递推式：dp[j] = max(dp[j], dp[j - nums[i]] + nums[i]);
// 3.初始理论：让dp数组在递归公式的过程中取的最大的价值，而不是被初始值覆盖了。本题题目中 只包含正整数的非空数组，所以非0下标的元素初始化为0就可以了。
// 4.外层遍历物品，内层倒序遍历背包
// 5.推导
package main

func canPartition(nums []int) bool {
	//求容量target
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum%2 == 1 {
		return false
	}
	target := sum / 2
	//定义并初始化
	dp := make([]int, target+1)
	//递推
	for i := 0; i < len(nums); i++ {
		for j := target; j >= nums[i]; j-- {
			dp[j] = max(dp[j], dp[j-nums[i]]+nums[i])
		}
	}
	return dp[target] == target
}
func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
