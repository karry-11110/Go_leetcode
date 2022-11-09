package main

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}
	result := 0
	if len(nums) == 1 {
		return 1
	}
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
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
