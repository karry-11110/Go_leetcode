package main

func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	result := 1
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			dp[i] = dp[i-1] + 1
		}
		if dp[i] > result {
			result = dp[i]
		}
	}
	return result
}
