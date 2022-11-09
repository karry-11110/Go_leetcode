// 1.完全背包和01背包问题唯一不同的地方就是，每种物品有无限件。
// 2.01背包和完全背包唯一不同就是体现在遍历顺序上，所以本文就不去做动规五部曲了，我们直接针对遍历顺序经行分析！
// 3.我们知道01背包内嵌的循环是从大到小遍历，为了保证每个物品仅被添加一次,而完全背包的物品是可以添加多次的，所以要从小到大去遍历
// 4.01背包中二维dp数组的两个for遍历的先后循序是可以颠倒了，一维dp数组的两个for循环先后循序一定是先遍历物品，再遍历背包容量。
// 	在完全背包中，对于一维dp数组来说，其实两个for循环嵌套顺序同样无所谓！
// 	因为dp[j] 是根据 下标j之前所对应的dp[j]计算出来的。 只要保证下标j之前的dp[j]都是经过计算的就可以了。
// test_CompletePack1 先遍历物品, 在遍历背包
package main

import "fmt"

func test_CompletePack1(weight, value []int, bagWeight int) int {
	// 定义dp数组 和初始化
	dp := make([]int, bagWeight+1)
	// 遍历顺序
	for i := 0; i < len(weight); i++ {
		// 正序会多次添加 value[i]
		for j := weight[i]; j <= bagWeight; j++ {
			// 推导公式
			dp[j] = max(dp[j], dp[j-weight[i]]+value[i])
			// debug
			//fmt.Println(dp)
		}
	}
	return dp[bagWeight]
}

// test_CompletePack2 先遍历背包, 在遍历物品
func test_CompletePack2(weight, value []int, bagWeight int) int {
	// 定义dp数组 和初始化
	dp := make([]int, bagWeight+1)
	// 遍历顺序
	// j从0 开始
	for j := 0; j <= bagWeight; j++ {
		for i := 0; i < len(weight); i++ {
			if j >= weight[i] {
				// 推导公式
				dp[j] = max(dp[j], dp[j-weight[i]]+value[i])
			}
			// debug
			//fmt.Println(dp)
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
	price := []int{15, 20, 30}
	fmt.Println(test_CompletePack1(weight, price, 4))
	fmt.Println(test_CompletePack2(weight, price, 4))
}
