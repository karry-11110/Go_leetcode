package main

// 1. 定义：dp[j] 表示：填满j（包括j）这么大容积的包，有dp[j]种方法
// 2. 递推式（组合类）：dp[j] += dp[j - nums[i]]
// 3. 初始化： dp[0]=1
// 4. 遍历顺序：nums放在外循环，target在内循环，且内循环倒序
// 5. 推导
import (
	"math"
)

func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if math.Abs(float64(target)) > float64(sum) {
		return 0
	}
	if (target+sum)%2 == 1 {
		return 0
	}
	//求背包容量
	bagSize := (sum + target) / 2
	//定义数组
	dp := make([]int, bagSize+1)
	//初始化
	dp[0] = 1
	//递推式
	for i := 0; i < len(nums); i++ {
		for j := bagSize; j >= nums[i]; j-- {
			dp[j] += dp[j-nums[i]]
		}
	}
	return dp[bagSize]
}
