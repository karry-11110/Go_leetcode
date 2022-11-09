package main

//贪心
// func maxSubArray(nums []int) int {
// 	result := math.MinInt64
// 	count := 0
// 	for i := 0; i < len(nums); i++ {
// 		count += nums[i]
// 		if count > result {
// 			result = count
// 		}
// 		if count <= 0 {
// 			count = 0
// 		}
// 	}
// 	return result
// }

//dp
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0]
	result := dp[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		if dp[i] > result {
			result = dp[i]
		}
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
